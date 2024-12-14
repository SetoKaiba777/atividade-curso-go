package controller

import (
	"api-pedidos/adapter/api/handler"
	"api-pedidos/adapter/api/response"
	"api-pedidos/core/erros"
	"api-pedidos/core/usecase"
	"api-pedidos/core/usecase/input"
	"net/http"
)

type SearchByOrderIdController struct {
	uc usecase.SearchById
}

func NewSearchByOrderIdController(uc usecase.SearchById) *SearchByOrderIdController {
	return &SearchByOrderIdController{uc: uc}
}

func (c *SearchByOrderIdController) Execute(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("orderId")

	if userId == "" {
		handler.HandleError(w, erros.NewInvalidRequestErr())
		return
	}

	i := &input.FindByIdInput{Id: userId}
	ctx := r.Context()
	pedidos, err := c.uc.Execute(&ctx, i)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response.NewSucess(http.StatusOK, pedidos).Send(w)
}
