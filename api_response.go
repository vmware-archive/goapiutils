package goapiutils

import (
	"net/http"
)

type ApiResponse struct {
	Body []byte
	Code int
}

func NewApiResponse(body []byte) *ApiResponse {
	return &ApiResponse{Body: body, Code: http.StatusOK}
}

func NewApiResponseWithCode(body []byte, code int) *ApiResponse {
	return &ApiResponse{Body: body, Code: code}
}
