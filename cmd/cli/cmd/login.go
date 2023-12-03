package datum

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/cmd/cli/vars"
	"github.com/datumforge/datum/internal/auth"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	scopes = []string{"email", "profile"}
)

// loginCmd represents the command to login to datum
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a Datum account",
	Long:  "Log in to a Datum account via clerk authentication oauth2.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return login(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login(ctx context.Context) error {
	ep := oauth2.Endpoint{
		AuthURL:  vars.AuthURL,
		TokenURL: vars.TokenURL,
	}

	oauthConfig := &oauth2.Config{
		ClientID:    vars.OauthClientID,
		Scopes:      scopes,
		Endpoint:    ep,
		RedirectURL: vars.RedirectURL,
	}

	te := TokenExchange{
		URL:       vars.TokenURL,
		GrantType: "authorization_code",
		TokenType: "jwt",
	}

	sc := ServiceConfig{
		Name:     "clerk",
		URL:      vars.ClerkRootURL,
		Exchange: &te,
	}

	token, err := auth.AuthPKCE(oauthConfig, vars.Audience)
	if err != nil {
		return err
	}

	if sc.Exchange != nil {
		e := auth.NewExchanger(sc.Exchange.URL, sc.Exchange.GrantType, sc.Exchange.TokenType)

		token, err = e.Exchange(ctx, token)
		if err != nil {
			return err
		}
	}

	fmt.Printf("auth token successfully retrieved, %v.\n", token)

	return nil
}

// ServiceConfig stores the config options for a service
type ServiceConfig struct {
	Name     string         `json:"name" yaml:"name"`
	URL      string         `json:"url" yaml:"url"`
	OIDC     EndpointOIDC   `json:"oidc" yaml:"oidc"`
	Exchange *TokenExchange `json:"exchange" yaml:"exchange"`
}

// EndpointOIDC stores the OIDC information for an endpoint
type EndpointOIDC struct {
	Audience string   `json:"audience" yaml:"audience"`
	Issuer   string   `json:"issuer" yaml:"issuer"`
	ClientID string   `json:"clientID" yaml:"clientID"`
	Scopes   []string `json:"scopes" yaml:"scopes"`
}

// TokenExchange stores token exchange information for an endpoint
type TokenExchange struct {
	URL       string `json:"url" yaml:"url"`
	GrantType string `json:"grantType" yaml:"grantType"`
	TokenType string `json:"tokenType" yaml:"tokenType"`
}
