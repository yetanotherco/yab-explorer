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

	ginMode := os.Getenv("GIN_MODE")
	if strings.ToLower(ginMode) == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	orders := router.Group("/orders")
	{
		orders.GET("", init.OrderController.GetOrders)

		orders.GET("/:orderId", init.OrderController.GetOrder)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
