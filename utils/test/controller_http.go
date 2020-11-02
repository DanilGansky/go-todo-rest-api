package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

// NewRequest ...
func NewRequest(method, path string, body []byte) (*httptest.ResponseRecorder, *http.Request) {
	var r *http.Request
	if len(body) == 0 {
		r = httptest.NewRequest(method, path, nil)
	} else {
		r = httptest.NewRequest(method, path, bytes.NewReader(body))
	}

	w := httptest.NewRecorder()
	return w, r
}

// MakeRequest ...
func MakeRequest(path string, handler func(http.ResponseWriter, *http.Request), w *httptest.ResponseRecorder, r *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc(path, handler)
	router.ServeHTTP(w, r)
}
