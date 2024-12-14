package router

import (
	"api-pedidos/adapter/api/controller"

	"github.com/gin-gonic/gin"
)

type (
	GinRouter interface {
		SetAppHandlers()
		GetRouter() *gin.Engine
	}

	ginEngine struct {
		router        *gin.Engine
		saveC         *controller.SaveController
		updateC       *controller.UpdateStatusController
		findByIdC     *controller.SearchByOrderIdController
		findByUserIdC *controller.SearchByUserIdController
	}
)

func NewGinEngine(
	router *gin.Engine,
	saveC *controller.SaveController,
	updateC *controller.UpdateStatusController,
	findByIdC *controller.SearchByOrderIdController,
	findByUserIdC *controller.SearchByUserIdController,
) *ginEngine {
	return &ginEngine{router: router, saveC: saveC, updateC: updateC,
		findByIdC: findByIdC, findByUserIdC: findByUserIdC}
}

func (e *ginEngine) SetAppHandlers() {
	e.router.GET("/v1/pedido/usuario/:userId", e.getByUserId())
	e.router.GET("/v1/pedido/:orderId", e.getByOrderId())
	e.router.POST("/v1/pedido", e.save())
	e.router.PUT("/v1/pedido", e.update())
}

func (e *ginEngine) getByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("userId", ctx.Param("userId"))
		ctx.Request.URL.RawQuery = query.Encode()
		e.findByUserIdC.Execute(ctx.Writer, ctx.Request)
	}
}

func (e *ginEngine) getByOrderId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("orderId", ctx.Param("orderId"))
		ctx.Request.URL.RawQuery = query.Encode()
		e.findByIdC.Execute(ctx.Writer, ctx.Request)
	}
}

func (e *ginEngine) save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.saveC.Execute(ctx.Writer, ctx.Request)
	}
}

func (e *ginEngine) update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.updateC.Execute(ctx.Writer, ctx.Request)
	}
}

func (e *ginEngine) GetRouter() *gin.Engine {
	return e.router
}
