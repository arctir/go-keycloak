// This file is *not* generated. These are convenience functions to compensate
// for a lacking API spec for posting realms to keycloak.
package keycloak

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type PostRealmJSONRequestBody = RealmRepresentation

// NewPostRealmRequest posts the keycloak api with an json realm representation
func NewPostRealmRequest(server string, body PostRealmJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostRequestWithBody(server, "application/json", bodyReader)
}

// PostRealm creates a new realm
func (c *Client) PostRealm(ctx context.Context, body PostRealmJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostRealmRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
