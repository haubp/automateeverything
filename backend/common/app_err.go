package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int   `json:"status_code"`
	RootErr    error `json:"-"`
	// user
	Message string `json:"message"`
	// system
	Log string `json:"log"`
	// custom key
	Key string `json:"error_key"`
}

// normal error, bad request
func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// Full
func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// 401, Unauthorrize
func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

// any type have function Error, we call that is error
// because this type is implement function Error of interface error
// so instead of returning type (*AppError), we can return error
func (e *AppError) Error() string {
	// return string of root error
	return e.RootError().Error()
}

func (e *AppError) RootError() error {
	// find root error
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	// found rootErr
	return e.RootErr
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "internal error", err.Error(), "INTERNAL_ERROR")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted ", strings.ToLower(entity)),
		fmt.Sprintf("Err %s Deleted", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot get %s ", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet %s", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s existed", strings.ToLower(entity)),
		fmt.Sprintf("%s Existed", entity),
	)
}

func ErrEntityNotExist(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not exist", strings.ToLower(entity)),
		fmt.Sprintf("%s Not exist", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot create %s ", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate %s", entity),
	)
}

func ErrInvalidLogin(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("email or password invalid"),
		fmt.Sprintf("ErrInvalidLogin"),
	)
}

func Err(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ErrNoPermission"),
	)
}

func GenerateJWTFail(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("genarate JWT fail"),
		fmt.Sprintf("GenerateJWTFail"),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ErrNoPermission"),
	)
}

func ErrInvalidPassword(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("You enter invalid password"),
		fmt.Sprintf("ErrInvalidPassword"),
	)
}

func CannotLogin(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not Login"),
		fmt.Sprintf("Generate Token fail"),
	)
}
