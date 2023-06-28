# go-keycloak

An auto-generated package for interacting with the Keycloak Admin API. This package is built from the OpenAPI specifications for the Keycloak API.

## Usage

```
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/arctir/go-keycloak"
	"golang.org/x/oauth2"
)

type passwordTokenSource struct {
	config   *oauth2.Config
	username string
	password string
}

func (s passwordTokenSource) Token() (*oauth2.Token, error) {
	return s.config.PasswordCredentialsToken(context.TODO(), s.username, s.password)
}

func main() {
	config := &oauth2.Config{
		ClientID: "admin-cli",
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:8080/realms/master/protocol/openid-connect/token",
		},
		Scopes: []string{"offline_access"}, // offline_access will provide a refresh token if enabled
	}

	oauth2client := oauth2.NewClient(context.TODO(), passwordTokenSource{
		config:   config,
		username: "admin",
		password: "password",
	})

	client, _ := keycloak.NewClient("http://localhost:8080/admin/realms", keycloak.WithHTTPClient(oauth2client))
	resp, err := client.GetRealm(context.TODO(), "foobar")
	...
}
```