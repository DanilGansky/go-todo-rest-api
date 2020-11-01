package response

import (
	"encoding/json"
	"net/http"
)

// HTTPError ...
type HTTPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// NewHTTPError ...
func NewHTTPError(status int, message string) *HTTPError {
	return &HTTPError{
		Status:  status,
		Message: message,
	}
}

// SendResponse ...
func SendResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	if status != 0 {
		w.WriteHeader(status)
	}

	json.NewEncoder(w).Encode(data)
}

// SendErrorResponse ...
func SendErrorResponse(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	httpErr := NewHTTPError(status, err.Error())
	httpErrJSON, _ := json.Marshal(httpErr)
	w.Write([]byte(httpErrJSON))
}
