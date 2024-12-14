package setup

import (
	"api-pedidos/adapter/api/controller"
	"api-pedidos/adapter/database"
	"api-pedidos/core/usecase"
	"api-pedidos/infrastructure/config"
	infraDb "api-pedidos/infrastructure/database"
	"api-pedidos/infrastructure/http/router"
	"api-pedidos/infrastructure/http/server"
	"api-pedidos/infrastructure/logger"
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type configuration struct {
	configApp *config.AppConfig
	webServer server.Server
	db        *database.PedidosRepository
	router    router.GinRouter
}

func NewConfig() *configuration {
	return &configuration{}
}

func (c *configuration) InitLogger() *configuration {
	logger.NewZapLogger()
	
	logger.Info("Log has been successfully configured")
	return c
}

func (c *configuration) WithAppConfig() *configuration {
	var err error
	c.configApp, err = config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	return c
}

func (c *configuration) WithDB() *configuration {
	db, err := infraDb.NewSQLConnection(c.configApp.MySQL.Host)
	if err != nil {
		logger.Fatal(err)
	}

	c.db = database.NewPedidosRepository(db)
	logger.Info("DB has been successfully configured")
	return c
}

func (c *configuration) WithRouter() *configuration {
	sc := controller.NewSaveController(usecase.NewSaveData(c.db))
	uc := controller.NewUpdateStatusController(usecase.NewUpdateStatus(c.db))
	foc := controller.NewSearchByOrderIdController(usecase.NewSearchById(c.db))
	fuc := controller.NewSearchByUserIdController(usecase.NewSearchByUserId(c.db))

	c.router = router.NewGinEngine(gin.Default(), sc, uc, foc, fuc)
	return c
}

func (c *configuration) WithWebServer() *configuration {
	intPort, err := strconv.ParseInt(c.configApp.Application.Server.Port, 10, 64)
	if err != nil {
		logger.Fatal(err)
	}

	intDuration, err := time.ParseDuration(c.configApp.Application.Server.Timeout + "s")
	if err != nil {
		logger.Fatal(err)
	}

	c.webServer = server.NewWebServer(c.router, intPort, intDuration*time.Second)
	logger.Info("Web server has been successfully configurated")
	return c
}

func (c *configuration) Start(ctx context.Context, wg *sync.WaitGroup) {
	logger.Info("App running on port %s", c.configApp.Application.Server.Port)
	c.webServer.Listen(ctx, wg)

}
