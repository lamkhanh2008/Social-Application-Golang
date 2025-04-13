package common

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewFullErrorResponse(status_code int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: status_code,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnAuthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

// func ErrorCan

func NewCustomError(root error, msg, key string) *AppError {
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

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something wrong in DB", err.Error(), "DB_Error")
}

func ErrCannotGetEntity(entityname string, err error) *AppError {
	return NewFullErrorResponse(http.StatusNotFound, err, fmt.Sprintf("Record %+v not found", entityname), err.Error(), "entity not found")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Error internal", err.Error(), "error internal")
}

func ErrInvalidRequest(err error) *AppError {
	return NewCustomError(err, err.Error(), "Invalid Request")
}

func RecordNotFound(entityName string, err error) *AppError {
	return NewFullErrorResponse(http.StatusNotFound, err, fmt.Sprintf("record of %s not found", entityName), err.Error(), "Record not found")
}

func ErrCannotCreateEntity(entityName string, err error) *AppError {
	return NewFullErrorResponse(http.StatusBadRequest, err, fmt.Sprintf("Create record of %s error", entityName), err.Error(), "Create Record err")
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(err, err.Error(), "ErrNoPermission")
}
