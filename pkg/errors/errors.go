package errors

import "net/http"

// ErrorResponse is the response for an error.
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

// Error is required by error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode returns the HTTP status code for the error.
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string) ErrorResponse {
	if msg == "" {
		msg = "Something went wrong. Please try again later."
	}

	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) ErrorResponse {
	if msg == "" {
		msg = "The request was invalid or cannot be otherwise served."
	}

	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

// Unauthorized creates a new error response representing an unauthorized request (HTTP 401)
func Unauthorized(msg string) ErrorResponse {
	if msg == "" {
		msg = "Unauthorized"
	}

	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: msg,
	}
}