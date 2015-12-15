package goapiutils

import "net/http"

type ApiHandler interface {
	Handle(req *http.Request) (*ApiResponse, *ApiError)
}
