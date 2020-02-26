package ce

import "net/http"

type Error struct {
	Code    int
	Message string
}

func (err *Error) NotFound() bool {
	return err.Code == http.StatusNotFound
}

func (err *Error) Unauthorized() bool {
	return err.Code == http.StatusUnauthorized
}

func (err *Error) TooManyRequests() bool {
	return err.Code == http.StatusTooManyRequests
}

func (err *Error) BadRequest() bool {
	return err.Code == http.StatusBadRequest
}

func (err *Error) Error() string {
	if err.Message == "" {
		return http.StatusText(err.Code)
	}

	return err.Message
}

func NewNotFound(message string) *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewValidationError(message string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewTooManyRequests(message string) *Error {
	return &Error{
		Code:    http.StatusTooManyRequests,
		Message: message,
	}
}

func NewBadRequest(message string) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorized(message string) *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewForbidden(message string) error {
	return &Error{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func BadRequestError(err error) bool {
	if err == nil {
		return false
	}

	if se, ok := err.(*Error); ok && se.BadRequest() {
		return true
	}

	return false
}

func TooManyRequestsError(err error) bool {
	if err == nil {
		return false
	}

	if se, ok := err.(*Error); ok && se.TooManyRequests() {
		return true
	}

	return false
}

func NotFoundError(err error) bool {
	if err == nil {
		return false
	}

	if se, ok := err.(*Error); ok && se.NotFound() {
		return true
	}

	return false
}
