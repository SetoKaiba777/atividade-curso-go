package controller

import (
	"api-pedidos/adapter/api/handler"
	"api-pedidos/adapter/api/response"
	"api-pedidos/core/usecase"
	"api-pedidos/core/usecase/input"
	"encoding/json"
	"io"
	"net/http"
)

type UpdateStatusController struct {
	uc usecase.UpdateStatus
}

func NewUpdateStatusController(uc usecase.UpdateStatus) *UpdateStatusController {
	return &UpdateStatusController{uc: uc}
}

func (c *UpdateStatusController) Execute(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := io.ReadAll(r.Body)
	if err != nil{
		handler.HandleError(w,err)
	}
	
	var i *input.UpdateStatusInput
	if err := json.Unmarshal(jsonBody, i); err != nil{
		handler.HandleError(w, err)
		return
	}

	ctx := r.Context()
	pedido, err := c.uc.Execute(&ctx, i)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response.NewSucess(http.StatusOK, pedido)
}