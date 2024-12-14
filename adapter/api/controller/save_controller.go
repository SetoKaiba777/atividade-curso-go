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

type SaveController struct {
	uc usecase.SaveData
}

func NewSaveController(uc usecase.SaveData) *SaveController {
	return &SaveController{uc: uc}
}

func (c *SaveController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	var i input.SaveInput
	if err := json.Unmarshal(jsonBody, &i); err != nil {
		handler.HandleError(w, err)
		return
	}

	p, err := c.uc.Execute(&ctx, &i)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response.NewSucess(http.StatusOK, p).Send(w)
}
