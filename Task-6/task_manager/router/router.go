package router

import (
	"net/http"
	"strings"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/controllers"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/middleware"
	"github.com/gin-gonic/gin"
)

func TaskHandlers(r *gin.RouterGroup, ctrl controllers.TaskHandlers) {

	// Tasks Endpoints
	r.GET("", ctrl.GetAllTasks())
	r.GET(":id", ctrl.GetTaskById())

	r.POST("", middleware.AdminAuthMiddleware(), ctrl.CreateTask())
	r.PUT(":id", middleware.AdminAuthMiddleware(), ctrl.UpdateTask())
	r.DELETE(":id", middleware.AdminAuthMiddleware(), ctrl.DeleteTask())

}

func UserHandlers(r *gin.RouterGroup, ctrl controllers.UserHandlers) {

	// Users Endpoints
	r.PUT(":id", ctrl.UpdateUser())
	r.GET(":id", ctrl.GetUserByUsername())
	r.DELETE(":id", ctrl.DeleteUser())

	// Admin Middleware
	r.Use(middleware.AdminAuthMiddleware())

	r.GET("", ctrl.GetAllUsers())
	r.PUT("/promote/:id", ctrl.PromoteUser())
	r.POST("/promote/:id", ctrl.PromoteUser())
	r.GET("/promote/:id", ctrl.PromoteUser())

}

func AuthHandlers(r *gin.RouterGroup, ctrl controllers.AuthHandlers) {

	// Auth Endpoints
	r.POST("/register", ctrl.SignUp())
	r.POST("/login", ctrl.Login())

}

func NoRouteHandler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		// Check if the route starts with "/example"
		if strings.HasPrefix(c.Request.URL.Path, "/users") || strings.HasPrefix(c.Request.URL.Path, "/tasks") || strings.HasPrefix(c.Request.URL.Path, "/auth") {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "Method not allowed",
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Route not found",
			})
		}
	})
}
