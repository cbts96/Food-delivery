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
	Key        string `json:"error_key"`
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

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}
func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
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
	} //đệ quy lỗi,  avoid sensitive error, hidding it, just show for dev

	return e.RootErr
}
func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with DB, please contact Admin", err.Error(), "DB_Err")
}
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with server, please contact Admin", err.Error(), "ErrInternal")
}
func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid Request Received", err.Error(), "ErrInvalidRequest")
}

func ErrCantListEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cant list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCantList %s", entity),
	)
}
func ErrCantDelteEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cant delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCantDelete %s", entity),
	)
}
func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("The value: %s is already exist", strings.ToLower(entity)),
		fmt.Sprintf("Err%sExisted", entity),
	)
}
func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s Not Found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}
func ErrCreateEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s cant be created", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}
func ErrNoPermission(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("You dont have permission to reach this resources"),
		fmt.Sprintf("ErrNoPermission"),
	)
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}
