package main

import (
	"github.com/GrandOichii/todoapp-golang/api/controllers"
	repositories "github.com/GrandOichii/todoapp-golang/api/repositories/task"
	services "github.com/GrandOichii/todoapp-golang/api/services/task"
	gin "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func createRouter() *gin.Engine {
	result := gin.Default()

	return result
}

func main() {
	config, err := readConfig("config.json")

	if err != nil {
		panic(err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	router := createRouter()

	taskController := controllers.CreateTaskController(
		services.CreateTaskServiceImpl(
			&repositories.TaskRepositoryImpl{},
			validate,
		),
	)

	taskController.Configure(router)

	router.Run("localhost:" + config.Port)
}
