package handler

import (
	"api-pedidos/adapter/api/response"
	"api-pedidos/core/erros"
	"errors"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	var status int

	switch {
	case errors.As(err, &erros.ChangeStatusErr{}),
		errors.As(err, &erros.InvalidRequestErr{}):
		status = http.StatusBadRequest
	case errors.As(err, &erros.NotFoundErr{}):
		status = http.StatusNotFound
	default:
		status = http.StatusInternalServerError
	}
	response.NewError(status, err).Send(w)
}
