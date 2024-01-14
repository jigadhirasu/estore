package types

import "encoding/json"

type ErrorImpl struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Error(code, message string, args ...any) string {

	err := ErrorImpl{
		Code:    code,
		Message: message,
	}

	b, _ := json.Marshal(err)
	return string(b)
}
