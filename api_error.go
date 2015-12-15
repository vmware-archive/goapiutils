package goapiutils

import (
	"encoding/json"
	"net/http"
)

const (
	KeyError   = "error"
	KeyMessage = "message"
)

type ApiError struct {
	message        string
	httpStatusCode int
}

func NewApiError(message string, httpStatusCode int) *ApiError {
	return &ApiError{
		message:        message,
		httpStatusCode: httpStatusCode,
	}
}

func WriteError(rw http.ResponseWriter, err *ApiError) {
	res, _ := json.Marshal(err)
	rw.WriteHeader(err.httpStatusCode)
	rw.Write(res)
}

func (self *ApiError) Message() string {
	return self.message
}

func (self *ApiError) Code() int {
	return self.httpStatusCode
}

func (self *ApiError) MarshalJSON() ([]byte, error) {
	errorMap := buildErrorMap(self.message)
	return json.Marshal(errorMap)
}

func buildErrorMap(errorMessage string) map[string]interface{} {
	return map[string]interface{}{
		KeyError: map[string]interface{}{
			KeyMessage: errorMessage,
		},
	}
}
