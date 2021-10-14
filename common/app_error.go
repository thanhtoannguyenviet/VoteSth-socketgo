package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Msg        string `json:"msg"`
	Log        string `json:"Log"`
	Key        string `json:"error_key"`
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func NewErrorResponse(root error, msg string, log string, key string) *AppError {
	return &AppError{
		RootErr: root,
		Msg:     msg,
		Log:     log,
		Key:     key,
	}
}
func NewFullErrorResponse(statusCode int, root error, msg string, log string, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Msg:        msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg string, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Msg:        msg,
		Log:        log,
		Key:        key,
	}
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with db", err.Error(), "DB_ERRO")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "INVALIDREQUEST")
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "internal error", err.Error(), "INTERNAL_ERROR")
}
func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("Cannot get %s", entity),
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}
func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ErrNoPermission"),
	)
}