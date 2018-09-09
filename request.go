package cha

import (
	"bytes"
	"net/http"
	"path"
)

// newRequest ... Create *http.Client for Chatwork
func (c Client) newRequest(
	method, spath string,
	params map[string]string,
) (req *http.Request, err error) {

	// Assemble URL from path
	uri := c.BaseURL
	uri.Path = path.Join(uri.Path, spath)

	if method != http.MethodGet {
		req, err = http.NewRequest(method, uri.String(), bytes.NewBufferString(BuildBody(params).Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest(http.MethodGet, BuildURL(uri, params).String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// Add APIKey to HTTP header
	req.Header.Set("X-ChatWorkToken", c.APIKey)

	return req, nil
}
