package utils

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

func NewCustomError(code int, format string, args ...interface{}) error {
	return CustomError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
