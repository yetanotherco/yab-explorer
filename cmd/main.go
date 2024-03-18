package main

import (
	"os"
	"yab-explorer/configs"
	"yab-explorer/controllers"
	"yab-explorer/repository"
	"yab-explorer/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Failed to load .env file")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")
	dbPort := os.Getenv("POSTGRES_PORT")

	database := configs.NewDBConnection(dbHost, dbUser, dbPassword, dbName, dbPort)

	orderRepository := repository.NewOrderRepository(database)
	orderService := services.New(orderRepository)
	orderController := controllers.New(orderService)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	orders := router.Group("/orders")
	{
		orders.GET("/", orderController.GetOrders)

		orders.GET("/:orderID", orderController.GetOrder)
	}

	apiPort := os.Getenv("API_PORT")

	router.Run(":" + apiPort)
}
