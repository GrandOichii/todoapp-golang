package main

import (
	"context"
	"fmt"

	"github.com/GrandOichii/todoapp-golang/api/config"
	"github.com/GrandOichii/todoapp-golang/api/controllers"
	repositories "github.com/GrandOichii/todoapp-golang/api/repositories/task"
	services "github.com/GrandOichii/todoapp-golang/api/services/task"
	gin "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createRouter() *gin.Engine {
	result := gin.Default()

	return result
}

func dbConnect(config *config.Configuration) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Db.ConnectionUri).SetServerAPIOptions(serverAPI))
	fmt.Printf("config.Db.DbName: %v\n", config.Db.DbName)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	config, err := config.ReadConfig("config.json")

	if err != nil {
		panic(err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	dbClient, err := dbConnect(config)
	if err != nil {
		panic(err)
	}

	router := createRouter()

	taskController := controllers.CreateTaskController(
		services.CreateTaskServiceImpl(
			repositories.CreateTaskRepositoryImpl(dbClient, config),
			validate,
		),
	)

	taskController.Configure(router)

	router.Run("localhost:" + config.Port)
}
