package middleware

import (
	"net/http"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/models"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the User from the context
		user, _ := c.MustGet("user").(models.UserRes)

		// Check if the user is an admin
		if user.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "only admins are authorized to perform this action"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()

	}
}
