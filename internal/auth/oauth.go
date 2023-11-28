package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	go_github "github.com/google/go-github/v56/github"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"

	"github.com/datumforge/datum/internal/utils/viperconfig"
)

const (
	Google = "google"
	Github = "github"
)

var knownProviders = []string{Google, Github}

// NewOAuthConfig creates a new OAuth2 config for the given provider and whether the client is a CLI or web client
func NewOAuthConfig(provider string, cli bool) (*oauth2.Config, error) {
	var err error

	redirectURL := func(provider string, cli bool) string {
		base := viper.GetString(fmt.Sprintf("%s.redirect_uri", provider))
		if cli {
			return fmt.Sprintf("%s/cli", base)
		}

		return fmt.Sprintf("%s/web", base)
	}

	scopes := func(provider string) []string {
		if provider == Google {
			return []string{"profile", "email"}
		}

		return []string{"user:email", "repo", "read:packages", "write:packages", "workflow", "read:org"}
	}

	endpoint := func(provider string) oauth2.Endpoint {
		if provider == Google {
			return google.Endpoint
		}

		return github.Endpoint
	}

	if provider != Google && provider != Github {
		return nil, newReadError("provider", provider, err)
	}

	clientId, err := readFileOrConfig(fmt.Sprintf("%s.clientID", provider)) //nolint:stylecheck

	if err != nil {
		return nil, newReadError("clientID", provider, err)
	}

	clientSecret, err := readFileOrConfig(fmt.Sprintf("%s.client_secret", provider))
	if err != nil {
		return nil, newReadError("clientID", provider, err)
	}

	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL(provider, cli),
		Scopes:       scopes(provider),
		Endpoint:     endpoint(provider),
	}, nil
}

// readFileOrConfig prefers reading from configKey_file for mounting of a file when running in an environment like Kubernetes, but falls back to a viper string value if the file is not present
func readFileOrConfig(configKey string) (string, error) {
	fileKey := configKey + "_file"
	if viper.IsSet(fileKey) && viper.GetString(fileKey) != "" {
		filename := viper.GetString(fileKey)
		// filepath.Clean avoids a gosec warning on reading a file by name
		data, err := os.ReadFile(filepath.Clean(filename))
		if err != nil {
			return "", err
		}

		return string(data), nil
	}

	return viper.GetString(configKey), nil
}

// NewProviderHTTPClient creates a new http client for the given provider
func NewProviderHTTPClient(provider string) *http.Client {
	if provider == Github {
		hClient := &http.Client{
			Transport: &go_github.BasicAuthTransport{
				Username: viper.GetString(fmt.Sprintf("%s.clientID", provider)),
				Password: viper.GetString(fmt.Sprintf("%s.client_secret", provider)),
			},
		}

		return hClient
	}

	return nil
}

// DeleteAccessToken deletes the access token for a given provider
func DeleteAccessToken(ctx context.Context, provider string, token string) error {
	var err error

	hClient := NewProviderHTTPClient(provider)
	if hClient == nil {
		return newReadError("provider", provider, err)
	}

	client := go_github.NewClient(hClient)
	clientID := viper.GetString(fmt.Sprintf("%s.clientID", provider))
	_, err = client.Authorizations.Revoke(ctx, clientID, token)

	if err != nil {
		return err
	}

	return nil
}

// ValidateProviderToken validates the given token for the given provider
func ValidateProviderToken(_ context.Context, provider string, token string) error {
	var err error

	if provider == Github {
		// Create an OAuth2 token source with the PAT
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

		// Create an authenticated GitHub client
		oauth2Client := oauth2.NewClient(context.Background(), tokenSource)
		client := go_github.NewClient(oauth2Client)

		// Make a sample API request to check token validity
		_, _, err := client.Users.Get(context.Background(), "")
		if err != nil {
			return newReadError("token", token, err)
		}

		return nil
	}

	return newReadError("provider", provider, err)
}

// RegisterOAuthFlags registers client ID and secret file flags for all known
// providers.  We need to build a consistent registration pattern other than this... TODO FUTURE MATT HAHA EAT SHIT
func RegisterOAuthFlags(v *viper.Viper, flags *pflag.FlagSet) error {
	for _, provider := range knownProviders {
		idFileKey := fmt.Sprintf("%s.client_id_file", provider)
		idFileFlag := fmt.Sprintf("%s-client-id-file", provider)
		idFileDesc := fmt.Sprintf("File containing %s client ID", provider)
		secretFileKey := fmt.Sprintf("%s.client_secret_file", provider)
		secretFileFlag := fmt.Sprintf("%s-client-secret-file", provider)
		secretFileDesc := fmt.Sprintf("File containing %s client secret", provider)

		if err := viperconfig.BindConfigFlag(
			v, flags, idFileKey, idFileFlag, "", idFileDesc, flags.String); err != nil {
			return err
		}

		if err := viperconfig.BindConfigFlag(
			v, flags, secretFileKey, secretFileFlag, "", secretFileDesc, flags.String); err != nil {
			return err
		}
	}

	return nil
}
