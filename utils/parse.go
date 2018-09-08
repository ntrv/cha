package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	cw "github.com/griffin-stewie/go-chatwork"
)

// parseThrottle ... Show response message for throttling
func parseThrottle(resp *http.Response) error {
	limit, err := rateLimit(resp.Header)
	if err != nil {
		return err
	}
	return fmt.Errorf("Throttled, ResetTime is %v", limit.ResetTime)
}

// parseError ... Show response message for Chatwork error
func parseError(body io.Reader) error {
	var eres cw.ChatWorkError

	res, err := parseBody(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, &eres)
	if err != nil {
		return err
	}
	return fmt.Errorf(strings.Join(eres.Errors, ", "))
}

// parseBody ... Show response message for common response
func parseBody(body io.Reader) ([]byte, error) {
	res, err := ioutil.ReadAll(body)
	if err != nil {
		return []byte(``), err
	}
	return res, nil
}
