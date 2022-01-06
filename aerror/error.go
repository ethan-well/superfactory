package aerror

import (
	"fmt"
)

type Error interface {
	error
	Code() string
	Message() string
}

type aError struct {
	PError  error
	code    string `json:"code"`
	message string `json:"message"`
}

func (e *aError) Code() string {
	return e.code
}

func (e *aError) Message() string {
	return e.message
}

func (e *aError) Error() string {
	if e.PError == nil {
		return fmt.Sprintf("code: %s, message: %s", e.Code(), e.Message())
	}

	return fmt.Sprintf("%s, code: %s, message: %s", e.PError, e.Code(), e.Message())
}

func NewError(err error, code string, message string) *aError {
	return &aError{
		PError:  err,
		code:    code,
		message: message,
	}
}

func NewErrorf(err error, code string, format string, opt ...interface{}) *aError {
	return &aError{
		PError:  err,
		code:    code,
		message: fmt.Sprintf(format, opt...),
	}
}
