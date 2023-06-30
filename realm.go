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

type PostRealmResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostRealmResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostRealmResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ParsePostRealmResponse parses an HTTP response from a PostRealmWithResponse call
func ParsePostRealmResponse(rsp *http.Response) (*PostRealmResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostRealmResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

func (c *ClientWithResponses) PostRealmWithResponse(ctx context.Context, body PostRealmJSONRequestBody, reqEditors ...RequestEditorFn) (*PostRealmResponse, error) {
	stdClient := c.ClientInterface.(*Client)
	rsp, err := stdClient.PostRealm(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostRealmResponse(rsp)
}
