package goapiutils

import (
	"encoding/json"
	"net/http"
)

func inferStatusCode(data interface{}) int {
	code := http.StatusOK

	if err, isApiError := data.(*ApiError); isApiError {
		code = err.Code()
	}
	return code
}

func WriteJSONData(w http.ResponseWriter, data interface{}) {
	code := inferStatusCode(data)
	WriteJSONWithCode(w, data, code)
}

func WriteJSONWithHeader(w http.ResponseWriter, data interface{}, headers http.Header) {
	code := inferStatusCode(data)
	WriteJSON(w, data, code, headers)
}

func WriteJSONWithCode(w http.ResponseWriter, data interface{}, statusCode int) {
	WriteJSON(w, data, statusCode, nil)
}

func WriteJSON(w http.ResponseWriter, data interface{}, statusCode int, headers http.Header) {
	content, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for name, values := range headers {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(content)
}
