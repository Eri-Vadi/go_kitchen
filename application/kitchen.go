package application

import (
	"context"
	"net/http"

	"github.com/Eri-Vadi/go_kitchen/internal/http/controller"
	"github.com/Eri-Vadi/go_kitchen/internal/infrastracture/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IApp interface {
	Start()
	Shutdown(ctx context.Context)
}

type kitchenApp struct {
	server *http.Server
}

func CreateApp() IApp {
	appHandler := gin.New()

	ctrl := controller.NewKitchenController()

	ctrl.RegisterKitchenRoutes(appHandler)

	app := kitchenApp{
		server: &http.Server{
			Addr:    viper.GetString("kitchen_host"),
			Handler: appHandler,
		},
	}

	return &app
}

func (d *kitchenApp) initialize() {

}

func (d *kitchenApp) Start() {
	d.initialize()
	logger.LogMessage("Starting kitchen server")

	if err := d.server.ListenAndServe(); err != http.ErrServerClosed {
		logger.LogPanicF("Unexpected error while running server: %v", err)
	}
}

func (d *kitchenApp) Shutdown(ctx context.Context) {
	if err := d.server.Shutdown(ctx); err != nil {
		logger.LogPanicF("Unexpected error while closing server: %v", err)
	}
	logger.LogMessage("Server terminated successfully")
}
