package router

import (
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/controllers"
	"github.com/gin-gonic/gin"
)

func TaskHandlers(r *gin.Engine, ctrl controllers.TaskHandlers) {

	// Tasks Endpoints
	r.GET("/tasks", ctrl.GetAllTasks())
	r.GET("/tasks/:id", ctrl.GetTaskById())
	r.PUT("/tasks/:id", ctrl.UpdateTask())
	r.DELETE("/tasks/:id", ctrl.DeleteTask())
	r.POST("/tasks", ctrl.CreateTask())

}

func UserHandlers(r *gin.Engine, ctrl controllers.UserHandlers) {

	// User Endpoints
	r.GET("/user", ctrl.GetAllUsers())
	r.GET("/user/:id", ctrl.GetUserByUsername())
	r.PUT("/user/:id", ctrl.UpdateUser())
	r.DELETE("/user/:id", ctrl.DeleteUser())
	r.POST("/user", ctrl.CreateUser())
	r.POST("/login", ctrl.Login())

}
