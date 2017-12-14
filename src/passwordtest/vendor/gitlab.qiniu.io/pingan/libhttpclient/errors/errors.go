package errors

import "fmt"

type RequestConstructorError struct {
	Method string
	Url    string
	Err    error
}

func (err RequestConstructorError) Error() string {
	return fmt.Sprintf("Failed to construct http request to %s %s: %s", err.Method, err.Url, err.Err)
}

type JSONDecodeError struct {
	Origin []byte
	Err    error
}

func (err JSONDecodeError) Error() string {
	return fmt.Sprintf("Failed to parse JSON from '%s': %s", err.Origin, err.Err)
}

type BodyCloseError struct {
	Err error
}

func (err BodyCloseError) Error() string {
	return fmt.Sprintf("Failed to close request body: %s", err.Err)
}

type IOReadError struct {
	Err error
}

func (err IOReadError) Error() string {
	return fmt.Sprintf("Failed to read request body: %s", err.Err)
}

type HTTPStatusError struct {
	StatusCode int
	Body       []byte
}

func (err HTTPStatusError) Error() string {
	return fmt.Sprintf("HTTP Failed: %d", err.StatusCode)
}
