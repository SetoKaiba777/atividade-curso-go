package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	StatusCode int
	Erros      []string
}

func NewError(status int, err error) *Error {
	return &Error{StatusCode: status, Erros: []string{err.Error()}}
}

func (e *Error) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	if err := json.NewEncoder(w).Encode(e.Erros); err != nil {
		return
	}
}