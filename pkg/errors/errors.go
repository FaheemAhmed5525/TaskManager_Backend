package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Err     error  `json:"error"`
}

func (app *AppError) Error() string {
	if app.Err != nil {
		fmt.Sprintf("%s: %s", app.Message, app.Error)
	}
	return app.Message
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func WrapError(err error, code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Commonly errors
var (
	ErrNotFound     = NewAppError(http.StatusNotFound, "resource not found")
	ErrUnauthorized = NewAppError(http.StatusUnauthorized, "unauthorized")
	ErrForBidden    = NewAppError(http.StatusForbidden, "forbidden")
	ErrBadRequest   = NewAppError(http.StatusBadRequest, "bad request")
	ErrInterServer  = NewAppError(http.StatusInternalServerError, "internal server error")
)

func ErrDatabase(err error) *AppError {
	return WrapError(err, http.StatusInternalServerError, "database error")
}

func ErrAlreadyExists() *AppError {
	return NewAppError(http.StatusConflict, "user already exists")
}

func ErrInvalidCredentials() *AppError {
	return NewAppError(http.StatusUnauthorized, "invalid credentials")
}
