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
	init := configs.Init()
	app := router.Init(init)

	app.Run(":" + apiPort)
}
