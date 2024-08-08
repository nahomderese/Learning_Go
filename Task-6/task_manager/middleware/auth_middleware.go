package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware(UserRepo data.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the Token from the request
		tokenString := c.GetHeader("Authorization")

		// Remove the "Bearer " prefix from the token
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if time.Now().Unix() > int64(claims["exp"].(float64)) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			username := claims["username"].(string)

			// Get the user from the database
			user, exist := UserRepo.FindByUsername(username)

			if !exist {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return
			}
			// Set the user in the context
			c.Set("user", user)

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()

	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the User from the context
		user, _ := c.MustGet("user").(models.User)

		// Check if the user is an admin
		if user.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "only admins are authorized to perform this action"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()

	}
}
