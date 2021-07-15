package model

import "fmt"

// HTTPError is a custom error to hold information regarding the HTTP status code.
type HTTPError struct {
	StatusCode int
	Message    string
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("Error code %d: %s", he.StatusCode, he.Message)
}
