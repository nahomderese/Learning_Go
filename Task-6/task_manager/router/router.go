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

	r.POST("", middleware.AdminAuthMiddleware(), ctrl.CreateTask())
	r.PUT(":id", middleware.AdminAuthMiddleware(), ctrl.UpdateTask())
	r.DELETE(":id", middleware.AdminAuthMiddleware(), ctrl.DeleteTask())

}

func UserHandlers(r *gin.RouterGroup, ctrl controllers.UserHandlers) {

	// Users Endpoints
	r.PUT(":username", ctrl.UpdateUser())
	r.GET(":username", ctrl.GetUser())
	r.DELETE(":username", ctrl.DeleteUser())

	r.GET("", middleware.AdminAuthMiddleware(), ctrl.GetAllUsers())
	r.PATCH("/promote/:username", middleware.AdminAuthMiddleware(), ctrl.PromoteUser())
	r.DELETE("", ctrl.DeleteAllUser())

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
