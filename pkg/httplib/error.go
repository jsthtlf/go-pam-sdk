package httplib

import (
	"encoding/json"
	"fmt"
)

var (
	CodeTerminalAlreadyExist = "terminal_already_exist"
	CodeObjectNotFound       = "object_does_not_exist"
	CodeAuthFailed           = "authentication_failed"
)

type ResponseError struct {
	Method  string
	UrlPath string
	Status  string

	Params map[string]interface{}
}

func (e *ResponseError) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.Params)
}

func (e *ResponseError) HasCode(code string) bool {
	if codeValue, ok := e.Params["code"]; ok {
		return codeValue == code
	}
	return false
}

func (e *ResponseError) Error() string {
	if len(e.Params) == 0 {
		return fmt.Sprintf("%s %s failed with %s", e.Method, e.UrlPath, e.Status)
	}

	params, _ := json.Marshal(e.Params)
	return fmt.Sprintf("%s %s failed with %s: %s)", e.Method, e.UrlPath, e.Status, string(params))
}
