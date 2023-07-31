package keycloak

import (
	"context"

	"golang.org/x/oauth2"
)

type clientCredentials struct {
	username string
	password string
	config   *oauth2.Config
}

type ClientCredentialsConfig struct {
	TokenURL string
	ClientID string
	Username string
	Password string
	Scopes   []string
}

func (c clientCredentials) Token() (*oauth2.Token, error) {
	token, err := c.config.PasswordCredentialsToken(context.Background(), c.username, c.password)
	return token, err
}

func (c ClientCredentialsConfig) NewClientCredentialsClientWithResponses(ctx context.Context, server string) (*ClientWithResponses, error) {

	clientScopes := append([]string{"offline_access"}, c.Scopes...)
	config := &oauth2.Config{
		ClientID: c.ClientID,
		Endpoint: oauth2.Endpoint{
			TokenURL: c.TokenURL,
		},
		Scopes: clientScopes,
	}

	oauth2client := oauth2.NewClient(ctx, clientCredentials{
		config:   config,
		username: c.Username,
		password: c.Password,
	})

	return NewClientWithResponses(server, WithHTTPClient(oauth2client))
}
