package utils

import "net/http"

// writeRes ... Create Response String from http.Response
func writeRes(res http.Response) ([]byte, error) {
	switch res.StatusCode {

	// Throttling
	case http.StatusTooManyRequests:
		return []byte(``), parseThrottle(res)

	// No Error
	case http.StatusOK:
		return parsebody(res.body)
	}

	// Error
	return []byte(``), parseError(res.Body)
}
