package main

import (
	"context"
	"log"
	"os"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/Delivery/controllers"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/Delivery/router"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/infrastructure"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/repositories"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	user_collection := client.Database("taskManager").Collection("users")
	task_collection := client.Database("taskManager").Collection("tasks")

	if err != nil {
		panic(err)
	}

	defer disconnect(client)

	// Repositories
	taskRepo := repositories.NewTaskRepository(task_collection)
	userRepo := repositories.NewUserRepository(user_collection)

	// Usecases
	taskUsecase := usecase.NewTaskUseCase(taskRepo)
	userUsecase := usecase.NewUserUseCase(userRepo)

	// Handlers (Routers)
	taskHandler := controllers.NewTaskHandlers(taskUsecase, userUsecase)
	userHandler := controllers.NewUserHandlers(userUsecase)
	authHandler := controllers.NewAuthHandlers(userUsecase)

	// Routes
	taskRoutes := r.Group("tasks")
	userRoutes := r.Group("users")
	authRoutes := r.Group("auth")

	// middlewares
	taskRoutes.Use(infrastructure.AuthMiddleware(userRepo))
	userRoutes.Use(infrastructure.AuthMiddleware(userRepo))

	// Handlers
	router.TaskHandlers(taskRoutes, taskHandler)
	router.UserHandlers(userRoutes, userHandler)
	router.AuthHandlers(authRoutes, authHandler)

	r.Run("localhost:8000")
}

func disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
