package router

import (
	"os"
	"strings"
	"yab-explorer/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(init *configs.Initialization) *gin.Engine {

	setGinMode(os.Getenv("GIN_MODE"))

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	orders := router.Group("/orders")
	{
		orders.GET("", init.OrderController.GetOrders)

		orders.GET("/:orderId", init.OrderController.GetOrder)

		orders.GET("/count", init.OrderController.GetOrdersCount)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func setGinMode(ginMode string) {
	switch strings.ToLower(ginMode) {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
