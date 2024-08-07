package router

import (
	"net/http"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/controllers"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/middleware"
	"github.com/gin-gonic/gin"
)

func TaskHandlers(r *gin.RouterGroup, ctrl controllers.TaskHandlers) {

	// Tasks Endpoints
	r.GET("", ctrl.GetAllTasks())
	r.GET(":id", ctrl.GetTaskById())

	// Admin Middleware
	r.Use(middleware.AdminAuthMiddleware())

	r.POST("", ctrl.CreateTask())
	r.PUT(":id", ctrl.UpdateTask())
	r.DELETE(":id", ctrl.DeleteTask())

}

func UserHandlers(r *gin.RouterGroup, ctrl controllers.UserHandlers) {

	// Users Endpoints
	r.PUT(":id", ctrl.UpdateUser())
	r.GET(":id", ctrl.GetUser())
	r.DELETE(":id", ctrl.DeleteUser())

	// Admin Middleware
	r.Use(middleware.AdminAuthMiddleware())

	r.GET("", ctrl.GetAllUsers())
	r.PATCH("/promote/:id", ctrl.PromoteUser())

}

func AuthHandlers(r *gin.RouterGroup, ctrl controllers.UserHandlers) {

	// Auth Endpoints
	r.POST("/register", ctrl.SignUp())
	r.POST("/login", ctrl.Login())

}

func NoRouteHandler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "The requested route was not found"})
	})
}
