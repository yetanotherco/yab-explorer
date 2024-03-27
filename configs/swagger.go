package configs

import (
	docs "yab-explorer/docs"
)

func InitSwagger(apiPort string) {
	docs.SwaggerInfo.Title = "YAB Explorer API"
	docs.SwaggerInfo.Description = "YAB Explorer API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + apiPort
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
}
