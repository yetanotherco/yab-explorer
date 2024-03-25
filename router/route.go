package router

import (
	"yab-explorer/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(init *configs.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	orders := router.Group("/orders")
	{
		orders.GET("", init.OrderController.GetOrders)

		orders.GET("/:orderId", init.OrderController.GetOrder)
	}

	return router
}
