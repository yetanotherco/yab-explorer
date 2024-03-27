package main

import (
	"os"
	"yab-explorer/configs"
	"yab-explorer/router"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	configs.InitLog()
}

func main() {

	apiPort := os.Getenv("API_PORT")

	if apiPort == "" {
		apiPort = "8080"
	}

	configs.InitSwagger(apiPort)

	init := configs.Init()
	app := router.Init(init)

	app.Run(":" + apiPort)
}
