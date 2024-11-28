package response

import (
	"encoding/json"
	"net/http"
)

type Sucess struct {
	StatusCode int
	Result     any
}

func NewSucess(status int, res any) *Sucess {
	return &Sucess{StatusCode: status, Result: res}
}

func (s *Sucess) Send(w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(s.StatusCode)
	if err := json.NewEncoder(w).Encode(s.Result); err != nil{
		return
	}
}