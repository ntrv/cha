package utils

import "net/http"

// WriteRes ... Create Response String from http.Response
func WriteRes(res *http.Response) ([]byte, error) {
	switch res.StatusCode {

	// Throttling
	case http.StatusTooManyRequests:
		return []byte(``), parseThrottle(res)

	// No Error
	case http.StatusOK:
		return parseBody(res.Body)
	}

	// Error
	return []byte(``), parseError(res.Body)
}
