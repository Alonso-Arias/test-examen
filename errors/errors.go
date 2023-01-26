package errors

import (
	"fmt"
)

// CustomError custom error
type CustomError struct {
	Message      string `json:"message"`
	Code         int    `json:"code"`
	InternalCode string `json:"internalCode"`
}

func (e CustomError) SetMessage(msg string) CustomError {
	e.Message = msg
	return e
}
func (e CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

var (
	BadRequest     = CustomError{Message: "BadRequest", Code: 400, InternalCode: "BADREQUEST"}
	Unauthorized   = CustomError{Message: "Unauthorized", Code: 401, InternalCode: "UNAUTHORIZED"}
	BlockedAccount = CustomError{Message: "Your account has been blocked", Code: 403, InternalCode: "BLOCKED_ACCOUNT"}
	NotFound       = CustomError{Message: "NotFound", Code: 403, InternalCode: "NOT_FOUND"}
	UpdateFailed   = CustomError{Message: "The record could not be updated", Code: 500, InternalCode: "UPDATE_FAILED"}
	InternalError  = CustomError{Message: "Error", Code: 500, InternalCode: "INTERNAL_SERVER_ERROR"}
)
