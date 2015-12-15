package goapiutils

import (
	"net/http"
)

var ErrInvalidRequest = NewApiError("Route not found", http.StatusNotFound)

type NotFoundHandler struct {
}

func NewNotFoundHandler() http.Handler {
	return &NotFoundHandler{}
}

func (self *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	WriteJSONData(w, ErrInvalidRequest)
}
