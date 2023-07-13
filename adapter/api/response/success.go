package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	statusCode int
	result     interface{}
}

func NewSuccess(result interface{}, statusCode int) Success {
	return Success{
		statusCode: statusCode,
		result:     result,
	}
}

func (s Success) Send(w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(s.statusCode)
	if errEncoder := json.NewEncoder(w).Encode(s.result); errEncoder != nil {
		return
	}
}