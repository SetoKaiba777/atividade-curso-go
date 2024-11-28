package controller

import (
	"api-pedidos/adapter/api/handler"
	"api-pedidos/adapter/api/response"
	"api-pedidos/core/erros"
	"api-pedidos/core/usecase"
	"api-pedidos/core/usecase/input"
	"net/http"
)

type SearchByUserIdController struct {
	uc usecase.SearchByUserId
}

func NewSearchByUserIdController(uc usecase.SearchByUserId) *SearchByUserIdController{
	return &SearchByUserIdController{uc: uc}
}

func (c *SearchByUserIdController) Execute(w http.ResponseWriter,r *http.Request ){
 	userId := r.URL.Query().Get("userId")
	
	if userId == ""{
		handler.HandleError(w, erros.NewInvalidRequestErr())
		return	
	}

	i := &input.FindByIdInput{Id: userId}
	ctx := r.Context()
	pedidos, err := c.uc.Execute(&ctx, i)
	if err != nil{
		handler.HandleError(w, err)
		return
	}

	response.NewSucess(http.StatusOK, pedidos)
}