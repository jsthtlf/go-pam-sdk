package httplib

import "fmt"

var (
	CodeTerminalAlreadyExist = "terminal_already_exist"
	CodeObjectNotFound       = "object_does_not_exist"
	CodeAuthFailed           = "authentication_failed"
)

type ErrResponseType struct {
	Method  string
	UrlPath string

	Detail string `json:"detail"`
	Code   string `json:"code"`
}

func (e ErrResponseType) Error() string {
	return fmt.Sprintf("%s %s failed: %s (Code: %s)", e.Method, e.UrlPath, e.Detail, e.Code)
}
