package main

import (
	"context"
	"log"
	"os"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/controllers"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/middleware"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())
	// r.Use(contentType)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	user_collection := client.Database("taskManager").Collection("users")
	task_collection := client.Database("taskManager").Collection("tasks")

	if err != nil {
		panic(err)
	}

	defer disconnect(err, client)

	taskRepo := data.NewMongoTaskRepository(task_collection)
	userRepo := data.NewMongoUserRepository(user_collection)

	taskHandler := &controllers.TaskController{TaskRepo: taskRepo, UserRepo: userRepo}
	userHandler := &controllers.UserController{UserRepo: userRepo}
	authHandler := &controllers.AuthController{UserRepo: userRepo}

	taskRoutes := r.Group("tasks")
	userRoutes := r.Group("users")
	authRoutes := r.Group("auth")

	// middlewares
	taskRoutes.Use(middleware.AuthMiddleware(userRepo))
	userRoutes.Use(middleware.AuthMiddleware(userRepo))

	router.TaskHandlers(taskRoutes, taskHandler)
	router.UserHandlers(userRoutes, userHandler)
	router.AuthHandlers(authRoutes, authHandler)

	r.Run("localhost:8000")
}

func disconnect(err error, client *mongo.Client) {
	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
