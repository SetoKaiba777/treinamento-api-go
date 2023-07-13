package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	StatusCode int
	Errors     []string
}

func NewError(err error, status int) *Error{
	return &Error{
		Errors: []string{err.Error()},
		StatusCode: status,
	}
}

func (e Error) Send(w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(e.StatusCode)
	if errEncoder := json.NewEncoder(w).Encode(e); errEncoder != nil {
		return
	}
}