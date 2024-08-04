package router

import (
	"github.com/Nahom-Derese/Learning_Go/Task-5/task_manager/controllers"
	"github.com/gin-gonic/gin"
)

func Handlers(r *gin.Engine, ctrl controllers.TaskHandlers) {
	r.GET("/tasks", ctrl.GetAllTasks())
	r.GET("/tasks/:id", ctrl.GetTaskById())
	r.PUT("/tasks/:id", ctrl.UpdateTask())
	r.DELETE("/tasks/:id", ctrl.DeleteTask())
	r.POST("/tasks", ctrl.CreateTask())
}
