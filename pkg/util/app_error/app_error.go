package app_error

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type AppError struct {
	StatusCode  int
	VisibleMsg  string
	InternalMsg string
	StackTrace  string
	Err         error
}

func (e *AppError) Error() string {
	var errStr string
	if e.Err != nil {
		errStr = e.Err.Error()
	} else {
		errStr = ""
	}
	return fmt.Sprintf("app error: code[%d], message[%s] user-message[%s] raw[%s] on[%s]", e.StatusCode, e.VisibleMsg, e.InternalMsg, errStr, e.StackTrace)
}

func NewAppErr(code int, visibleMsg string, internalMsg string, err error) *AppError {
	stack := zap.Stack("").String
	return &AppError{
		StatusCode:  code,
		VisibleMsg:  visibleMsg,
		InternalMsg: internalMsg,
		StackTrace:  stack,
		Err:         err,
	}
}

func IsAppError(err error) (*AppError, bool) {
	var me *AppError
	ok := errors.As(err, &me)
	return me, ok
}
