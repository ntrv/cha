package utils

import (
	"net/url"
	"strings"
)

func buildURL(endpoint *url.URL, params map[string]string) *url.URL {
	var query []string

	for k := range params {
		query = append(query, k+"="+params[k])
	}

	ep := *endpoint
	ep.RawQuery = strings.Join(query, "&")
	return &ep
}

func buildBody(params map[string]string) url.Values {
	body := url.Values{}
	for k := range params {
		body.Add(k, params[k])
	}
	return body
}
