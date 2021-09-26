package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Eri-Vadi/go_kitchen/internal/domain/dto"
	"github.com/Eri-Vadi/go_kitchen/internal/domain/repository"
	"github.com/Eri-Vadi/go_kitchen/internal/http/httperr"
	"github.com/Eri-Vadi/go_kitchen/internal/infrastracture/logger"
	"github.com/Eri-Vadi/go_kitchen/internal/service/supervisor"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const CurrentCaller = "Kitchen Controller"

type IController interface {
	menu(c *gin.Context)
	order(c *gin.Context)
	RegisterKitchenRoutes(c *gin.Engine)
}

type KitchenController struct {
	super supervisor.ISupervisor
}

func NewKitchenController() IController {
	return &KitchenController{
		super: supervisor.NewKitchenSupervisor(),
	}
}

func (ctrl *KitchenController) menu(c *gin.Context) {
	var response dto.Menu

	response.Items = repository.GetFoods()
	response.ItemsCount = len(response.Items)

	c.JSON(http.StatusOK, response)
}

func (ctrl *KitchenController) order(c *gin.Context) {
	var currentOrder dto.Order

	if err := c.ShouldBindJSON(&currentOrder); err != nil {
		httperr.HandleErr(CurrentCaller, err, c)
		return
	}

	log.Printf("%+v", currentOrder)
	ctrl.super.PrepareOrder(currentOrder)

	logger.LogMessageF("Order %v completed", currentOrder.OrderID)

	resp := dto.Distribution{}
	resp.TableID = currentOrder.TableID

	jsonBody, err := json.Marshal(resp)
	if err != nil {
		log.Panic(err)
	}
	contentType := "application/json"

	http.Post(viper.GetString("dining_host")+"/distribution", contentType, bytes.NewReader(jsonBody))

	return
}
