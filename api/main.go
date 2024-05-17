package main

import (
	"github.com/GrandOichii/todoapp-golang/api/config"
	docs "github.com/GrandOichii/todoapp-golang/api/docs"
	"github.com/GrandOichii/todoapp-golang/api/router"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title TODOapp api
// @version 1.0
// @description A siple TODO task service
// @host localhost:8080
// @BasePath /api/v1
func main() {
	config, err := config.ReadConfig("config.json")

	if err != nil {
		panic(err)
	}

	router := router.CreateRouter(config)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":" + config.Port)
}
