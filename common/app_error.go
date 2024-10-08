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
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

// convert AppError to Error implement Error() String
func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func NewErrorResponse(rootError error, msg, log, key string) *AppError {
	return &AppError{StatusCode: http.StatusBadRequest, RootErr: rootError, Message: msg, Log: log, Key: key}
}

func NewFullErrorResponse(statusCode int, rootError error, msg, log, key string) *AppError {
	return &AppError{StatusCode: statusCode, RootErr: rootError, Message: msg, Log: log, Key: key}
}

func NewUnauthenticateResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func ErrDb(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with db", err.Error(), "DB_SERVER_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ERR_INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong in the server",
		err.Error(),
		"ErrInternal",
	)
}

func ErrPermission() *AppError {
	return NewFullErrorResponse(
		403,
		errors.New("permission denied"),
		"Permission denied",
		"Permission denied",
		"ErrPermissionDenied",
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

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create record %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrRecordNotFound(err error) *AppError {
	return NewFullErrorResponse(400, err, "Record not found", err.Error(), "ErrRecordNotFound")
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Update record %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}
