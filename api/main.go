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
// @host localhost:9090
// @BasePath /api/v1
func main() {
	var c *config.Configuration

	c, err := config.ReadConfig("config.json")

	if err != nil {
		c, err = config.ReadEnvConfig()
		if err != nil {
			panic(err)
		}
	}

	router := router.CreateRouter(c)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":" + c.Port)
}
