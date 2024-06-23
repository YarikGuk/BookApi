package error

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrInvalidInput     = NewAppError(http.StatusBadRequest, "Invalid input")
	ErrNotFound         = NewAppError(http.StatusNotFound, "Resource not found")
	ErrInternalServer   = NewAppError(http.StatusInternalServerError, "Internal server error")
	ErrFutureBirthDate  = NewAppError(http.StatusBadRequest, "Birth date cannot be in the future")
	ErrInvalidBirthDate = NewAppError(http.StatusBadRequest, "Invalid birth date format. Use yyyy-mm-dd")
)
