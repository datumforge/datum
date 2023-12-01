package datum

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/adrg/xdg"
	"github.com/alexkappa/mustache"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"

	"github.com/datumforge/datum/cmd/cli/vars"
)

// HTML is shown to the user after they log in.
var HTML = `<!doctype html>` + "\n" + //nolint:gochecknoglobals
	`<html>
        <head>
                <meta charset="utf-8"/>
                <title>Login</title>
                <!--[if lt IE 9]>
                <script src="//html5shim.googlecode.com/svn/trunk/html5.js"></script>
                <![endif]-->
        </head>
        <body>
            <h1>Successfully logged in as {{email}}</h1>
        <h2>Please close this tab or page and return to the CLI</h2>
        </body>
</html>`

const (
	ServerPort             = 17608
	ClerkClientSessionPath = "%s/v1/client?_clerk_js_version=4.50.1&__dev_session=%s"
)

// loginData is the data that is stored in the jwt.json file.
type loginData struct {
	UserID    string `json:"userId"`
	SessionID string `json:"sessionId"`
	Token     string `json:"token"`
}

type token struct {
	Jwt string `json:"jwt"`
}

type verification struct {
	Status string `json:"status"`
}

type email struct {
	ID           string       `json:"id"`
	Address      string       `json:"email_address"` //nolint:tagliatelle
	Verification verification `json:"verification"`
}

type phone struct {
	ID           string       `json:"id"`
	Number       string       `json:"phone_number"` //nolint:tagliatelle
	Verification verification `json:"verification"`
}

type user struct {
	ID             string  `json:"id"`
	Username       string  `json:"username"`
	FirstName      string  `json:"first_name"`               //nolint:tagliatelle
	LastName       string  `json:"last_name"`                //nolint:tagliatelle
	ImageURL       string  `json:"image_url"`                //nolint:tagliatelle
	PrimaryEmail   string  `json:"primary_email_address_id"` //nolint:tagliatelle
	PrimaryPhone   string  `json:"primary_phone_number_id"`  //nolint:tagliatelle
	EmailAddresses []email `json:"email_addresses"`          //nolint:tagliatelle
	PhoneNumbers   []phone `json:"phone_numbers"`            //nolint:tagliatelle
}

type session struct {
	LastActiveToken token `json:"last_active_token"` //nolint:tagliatelle
	CreatedAt       int64 `json:"created_at"`        //nolint:tagliatelle
	UpdatedAt       int64 `json:"updated_at"`        //nolint:tagliatelle
	User            user  `json:"user"`
}

type response struct {
	Sessions []session `json:"sessions"`
}

type clientResponse struct {
	Response response `json:"response"`
}

type handler struct{}

const WaitBeforeExitSeconds = 3

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// This path is followed after the user logs in. The CLI Auth Client redirects to here.
	switch {
	case request.URL.Path == "/done" && request.Method == "GET":
		bts, _ := base64.StdEncoding.DecodeString(request.URL.Query().Get("p"))

		rsp, loginEmail, err := processLogin(request.Context(), bts, true)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			log.Printf("error: %v", err)

			return
		}

		writer.WriteHeader(http.StatusOK)

		// nosemgrep: go.lang.security.audit.xss.no-direct-write-to-responsewriter.no-direct-write-to-responsewriter
		_, _ = writer.Write([]byte(rsp))

		go func() {
			// Tell the user we're done and then forcefully exit the program.
			fmt.Printf("Successfully logged in as %s\n", loginEmail)
			time.Sleep(WaitBeforeExitSeconds * time.Second)
			os.Exit(0)
		}()

		return
	case request.URL.Path == "/" && request.Method == "GET":
		writer.Header().Set("Location", vars.LoginURL)
		writer.WriteHeader(http.StatusTemporaryRedirect)
	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}

const JwtFilePermissions = 0o600

// processLogin takes the JWT token, verifies it, and then stores it in the jwt.json file.
func processLogin(ctx context.Context, payload []byte, write bool) (string, string, error) {
	path := getJwtPath()
	if write {
		if err := os.WriteFile(path, pretty.Pretty(payload), JwtFilePermissions); err != nil {
			return "", "", err
		}
	}

	data := &loginData{}
	if err := json.Unmarshal(payload, data); err != nil {
		return "", "", err
	}

	// Call out to clerk and ask for session info using the JWT token.
	u := fmt.Sprintf(ClerkClientSessionPath, vars.ClerkRootURL, data.Token)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Origin", fmt.Sprintf("http://localhost:%d", ServerPort))

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}

	defer func() { _ = rsp.Body.Close() }()

	bb, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", "", err
	}

	if rsp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("http %d", rsp.StatusCode) //nolint:goerr113
	}

	cr := &clientResponse{}
	if err := json.Unmarshal(bb, cr); err != nil {
		return "", "", err
	}

	return decodeJWT(cr.Response.Sessions[0].LastActiveToken.Jwt)
}

func decodeJWT(jwt string) (string, string, error) {
	// Using a dummy value here because DecodeToken doesn't actually use the secret.
	c, err := clerk.NewClient("dummy")
	if err != nil {
		return "", "", err
	}

	// Extract the claims (which includes the email address) from the JWT token.
	claims, err := c.DecodeToken(jwt)
	if err != nil {
		return "", "", err
	}

	// Grab the email address from the claims.
	emailStr, ok := claims.Extra["email"].(string)
	if !ok {
		return "", "", fmt.Errorf("couldn't find email address in claims") //nolint:goerr113
	}

	ht, err := getHTML(emailStr)
	if err != nil {
		return "", "", err
	}

	// Return the HTML and email
	return ht, emailStr, nil
}

func getHTML(emailStr string) (string, error) {
	// Render the HTML
	tmpl := mustache.New()
	if err := tmpl.ParseString(HTML); err != nil {
		return "", err
	}

	ht, err := tmpl.RenderString(map[string]string{
		"email": emailStr,
	})
	if err != nil {
		return "", err
	}

	return ht, nil
}

const ReadHeaderTimeoutSeconds = 3

// loginCmd represents the login command.
var loginCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "login",
	Short: "Log in to an datum account",
	Long:  "Log in to an datum account.",
	Run: func(cmd *cobra.Command, args []string) {
		needLogin := false
		path := getJwtPath()
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				needLogin = true
			} else {
				log.Fatalln(err)
			}
		}

		if needLogin {
			http.Handle("/", &handler{})
			go func() {
				time.Sleep(1 * time.Second)
				openBrowser(fmt.Sprintf("http://localhost:%d", ServerPort))
			}()

			server := &http.Server{
				Addr:              fmt.Sprintf(":%d", ServerPort),
				ReadHeaderTimeout: ReadHeaderTimeoutSeconds * time.Second,
			}

			// nosemgrep: go.lang.security.audit.net.use-tls.use-tls
			log.Fatalln(server.ListenAndServe())
		} else {
			if fileInfo.IsDir() {
				log.Fatalln("jwt path isn't a regular file:", path)
			}

			contents, err := os.ReadFile(path)
			if err != nil {
				log.Fatalln(err)
			}

			_, loginEmail, err := processLogin(cmd.Context(), contents, false)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("You're already logged in as %s\n", loginEmail) //nolint:forbidigo

			os.Exit(0)
		}
	},
}

func getJwtName() string {
	if vars.Stage == "local" {
		return "/Users/manderson/datum/jwt.json"
	}

	return fmt.Sprintf("datum/jwt-%s.json", vars.Stage)
}

// getJwtPath returns the path to the jwt.json file where the JWT token is stored.
func getJwtPath() string {
	path, err := xdg.ConfigFile(getJwtName())
	if err != nil {
		log.Fatalln(err)
	}

	return path
}

// openBrowser tries to open the URL in a browser. Should work on most standard platforms.
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform: %s", runtime.GOOS) //nolint:goerr113
	}

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
