package httplib

import (
	"encoding/json"
	"fmt"
)

var (
	CodeTerminalAlreadyExist         = "terminal_already_exist"
	CodeTerminalRegistrationDisabled = "terminal_registration_disabled"
	CodeObjectNotFound               = "object_does_not_exist"
	CodeAuthFailed                   = "authentication_failed"
	CodeAuthNotProvided              = "not_authenticated"
	CodeLicenseValidateError         = "license_validate_error"
	CodeLicenseLimitSessions         = "license_limit_sessions"
)

var detailMessages = map[string]string{
	CodeTerminalAlreadyExist:         "Terminal with this name already exists",
	CodeTerminalRegistrationDisabled: "Terminal registration is disabled",
	CodeObjectNotFound:               "Requested object does not exist",
	CodeAuthFailed:                   "The terminal cannot be authenticated",
	CodeAuthNotProvided:              "Authentication credentials were not provided",
	CodeLicenseValidateError:         "The license does not exist or cannot be validated",
	CodeLicenseLimitSessions:         "The license session limit has been exceeded",
}

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
	if codeValue, ok := e.Params["code"].(string); ok {
		return codeValue == code
	}
	return false
}

func (e *ResponseError) Error() string {
	if codeValue, ok := e.Params["code"].(string); ok {
		if msg, ok := detailMessages[codeValue]; ok {
			return msg
		}
	}

	if len(e.Params) == 0 {
		return fmt.Sprintf("%s %s failed with %s", e.Method, e.UrlPath, e.Status)
	}

	params, _ := json.Marshal(e.Params)
	return fmt.Sprintf("%s %s failed with %s: %s)", e.Method, e.UrlPath, e.Status, string(params))
}
