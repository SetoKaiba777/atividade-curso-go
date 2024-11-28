package handler

import (
	"api-pedidos/adapter/api/response"
	"api-pedidos/core/erros"
	"errors"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error){
	var status int

	switch {
	case errors.Is(erros.ChangeStatusErr{}, err):
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}
	response.NewError(status, err).Send(w)
}