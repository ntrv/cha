package client

import (
	"net/http"
	"path"
	"bytes"

	"github.com/ntrv/cha/client/utils"
)

// newRequest ... Create *http.Client for Chatwork
func (c Client) newRequest(
	method, spath string,
	params map[string]string,
) (*http.Request, err error) {

	// Assemble URL from path
	uri := c.BaseURL
	uri.Path = path.Join(uri.Path, spath)

	if method != http.MethodGet {
		req, err = http.NewRequest(method, uri.String(), bytes.NewBufferString(utils.buildBody(params).Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest(http.MethodGet, utils.buildURL(uri, params).String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// Add APIKey to HTTP header
	req.Header.Set("X-ChatWorkToken", c.APIKey)

	return req, nil
}


