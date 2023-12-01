package datum

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/datumforge/datum/cmd/cli/vars"
)

// tokenCmd represents the generate-request-token command.
var tokenCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "generate-request-token",
	Short: "Generate a request token",
	Long: "Generate a JWT token to be used for HTTP requests, and prints it." +
		" This command is useful for testing purposes.",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		contents, err := os.ReadFile(getJwtPath())
		if err != nil {
			log.Fatalln(err)
		}

		data := &loginData{}
		if err := json.Unmarshal(contents, data); err != nil {
			log.Fatalln(err)
		}

		// Call out to clerk and ask for session info using the JWT token.
		u := fmt.Sprintf(ClerkClientSessionPath, vars.ClerkRootURL, data.Token)
		req, err := http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			log.Fatalln(err)
		}

		rsp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		bb, err := io.ReadAll(rsp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		if rsp.StatusCode != http.StatusOK {
			log.Fatalln(fmt.Errorf("http %d", rsp.StatusCode)) //nolint:goerr113
		}

		cr := &clientResponse{}
		if err := json.Unmarshal(bb, cr); err != nil {
			log.Fatalln(err)
		}

		jwt := cr.Response.Sessions[0].LastActiveToken.Jwt
		fmt.Println(jwt) //nolint:forbidigo
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
