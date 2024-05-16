package router

import (
	"net/http"
	"time"

	"github.com/GrandOichii/todoapp-golang/api/config"
	"github.com/GrandOichii/todoapp-golang/api/controllers"
	"github.com/GrandOichii/todoapp-golang/api/middleware"
	"github.com/gin-contrib/cors"

	"context"

	taskrepositories "github.com/GrandOichii/todoapp-golang/api/repositories/task"
	userrepositories "github.com/GrandOichii/todoapp-golang/api/repositories/user"
	taskservices "github.com/GrandOichii/todoapp-golang/api/services/task"
	userservices "github.com/GrandOichii/todoapp-golang/api/services/user"
	gin "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRouter(config *config.Configuration) *gin.Engine {
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

	dbClient, err := dbConnect(config)
	if err != nil {
		panic(err)
	}

	// repositories
	userRepo := userrepositories.CreateUserRepositoryImpl(dbClient, config)
	taskRepo := taskrepositories.CreateTaskRepositoryImpl(dbClient, config)

	configRouter(result, userRepo, taskRepo)

	result.GET("/api/v1/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hi!")
	})

	return result
}

func configRouter(router *gin.Engine, userRepo userrepositories.UserRepository, taskRepo taskrepositories.TaskRepository) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	// services
	userService := userservices.CreateUserServiceImpl(
		userRepo,
		validate,
	)

	// middleware
	auth := middleware.CreateJwtMiddleware(userService)

	// controllers
	taskController := controllers.CreateTaskController(
		taskservices.CreateTaskServiceImpl(
			taskRepo,
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
}

func dbConnect(config *config.Configuration) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Db.ConnectionUri).SetServerAPIOptions(serverAPI))
	if err != nil {
		return nil, err
	}
	return client, nil
}
