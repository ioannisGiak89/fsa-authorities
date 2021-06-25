package model

import "fmt"

// error the error interface used to create custom errors
type error interface {
	Error() string
}

// HttpError sustom error to hold information regarding the HTTTP status code
type HttpError struct {
	StatusCode int
	Message    string
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("Error code %d: %s", he.StatusCode, he.Message)
}
