package main

import (
	"context"
	"log"
	"os"

	"github.com/Nahom-Derese/Learning_Go/Task-5/task_manager/controllers"
	"github.com/Nahom-Derese/Learning_Go/Task-5/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-5/task_manager/router"
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
	r := gin.New()
	r.Use(gin.Recovery())

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	collection := client.Database("taskManager").Collection("tasks")

	if err != nil {
		panic(err)
	}

	defer disconnect(err, client)

	taskRepo := data.NewMongoTaskRepository(collection)

	taskHandler := &controllers.TaskController{TaskRepo: taskRepo}

	router.Handlers(r, taskHandler)

	r.Run("localhost:8080")
}
func disconnect(err error, client *mongo.Client) {
	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
