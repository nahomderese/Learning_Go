package main

import (
	"github.com/Nahom-Derese/Learning_Go/Task-4/task_manager/controllers"
	"github.com/Nahom-Derese/Learning_Go/Task-4/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-4/task_manager/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(contentType)

	taskRepo := data.NewInMemoryTaskRepository()

	taskHandler := &controllers.TaskController{TaskRepo: taskRepo}

	router.Handlers(r, taskHandler)

	r.Run("localhost:8000")
}

func contentType(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
}
