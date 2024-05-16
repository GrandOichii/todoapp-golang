package main

import (
	"context"
	"time"

	"github.com/GrandOichii/todoapp-golang/api/config"
	"github.com/GrandOichii/todoapp-golang/api/controllers"
	"github.com/GrandOichii/todoapp-golang/api/middleware"
	taskrepositories "github.com/GrandOichii/todoapp-golang/api/repositories/task"
	userrepositories "github.com/GrandOichii/todoapp-golang/api/repositories/user"
	taskservices "github.com/GrandOichii/todoapp-golang/api/services/task"
	userservices "github.com/GrandOichii/todoapp-golang/api/services/user"
	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	docs "github.com/GrandOichii/todoapp-golang/api/docs"
	swaggerfiles "github.com/swaggo/files"
)

func createRouter() *gin.Engine {
	result := gin.Default()

	result.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			// return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	return result
}

func dbConnect(config *config.Configuration) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Db.ConnectionUri).SetServerAPIOptions(serverAPI))
	if err != nil {
		return nil, err
	}
	return client, nil
}

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

	validate := validator.New(validator.WithRequiredStructEnabled())

	dbClient, err := dbConnect(config)
	if err != nil {
		panic(err)
	}

	router := createRouter()

	// services
	userService := userservices.CreateUserServiceImpl(
		userrepositories.CreateUserRepositoryImpl(
			dbClient, config,
		),
		validate,
	)

	// middleware
	auth := middleware.CreateJwtMiddleware(userService)

	// controllers
	taskController := controllers.CreateTaskController(
		taskservices.CreateTaskServiceImpl(
			taskrepositories.CreateTaskRepositoryImpl(dbClient, config),
			validate,
		),
		auth.Middle.MiddlewareFunc(),
	)
	taskController.Configure(router)

	authController := controllers.CreateAuthController(
		userService,
		auth.Middle.LoginHandler,
	)
	authController.Configure(router)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:" + config.Port)
}
