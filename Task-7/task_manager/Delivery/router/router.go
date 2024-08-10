package router

import (
	"net/http"

	"github.com/Nahom-Derese/Learning_Go/Task-7/task-manager/Delivery/controllers"
	infrastructure "github.com/Nahom-Derese/Learning_Go/Task-7/task-manager/infrastructure"
	"github.com/gin-gonic/gin"
)

func TaskHandlers(r *gin.RouterGroup, ctrl controllers.TaskHandlers) {

	// Tasks Endpoints
	r.GET("", ctrl.GetAllTasks())
	r.GET(":id", ctrl.GetTaskById())

	r.Use(infrastructure.AdminAuthMiddleware())
	r.POST("", ctrl.CreateTask())
	r.PUT(":id", ctrl.UpdateTask())
	r.DELETE(":id", ctrl.DeleteTask())

}

func UserHandlers(r *gin.RouterGroup, ctrl controllers.UserHandlers) {

	// Users Endpoints
	r.PUT(":username", ctrl.UpdateUser())
	r.GET(":username", ctrl.GetUser())
	r.DELETE(":username", ctrl.DeleteUser())

	r.Use(infrastructure.AdminAuthMiddleware())
	r.GET("", ctrl.GetAllUsers())
	r.PATCH("/promote/:username", ctrl.PromoteUser())

}

func AuthHandlers(r *gin.RouterGroup, ctrl controllers.AuthHandlers) {

	// Auth Endpoints
	r.POST("/register", ctrl.SignUp())
	r.POST("/login", ctrl.Login())

}

func NoRouteHandler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "The requested route was not found"})
	})
}
