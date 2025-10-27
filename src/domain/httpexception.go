package domain

import "fmt"

// HTTPException represents an HTTP error.
// Pure domain exception, no external dependencies.
type HTTPException struct {
	StatusCode int
	Message    string
	Details    map[string]interface{}
}

// Error implements the error interface.
func (e *HTTPException) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

// NewHTTPException creates a new HTTP exception.
func NewHTTPException(statusCode int, message string) *HTTPException {
	return &HTTPException{
		StatusCode: statusCode,
		Message:    message,
		Details:    make(map[string]interface{}),
	}
}

// Common HTTP exceptions
var (
	BadRequest    = NewHTTPException(400, "Bad Request")
	Unauthorized  = NewHTTPException(401, "Unauthorized")
	Forbidden     = NewHTTPException(403, "Forbidden")
	NotFound      = NewHTTPException(404, "Not Found")
	Conflict      = NewHTTPException(409, "Conflict")
	InternalError = NewHTTPException(500, "Internal Server Error")
)

