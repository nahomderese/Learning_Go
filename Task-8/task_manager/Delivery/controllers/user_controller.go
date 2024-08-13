package controllers

import (
	"net/http"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/gin-gonic/gin"
)

type UserHandlers interface {
	GetAllUsers() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	PromoteUser() gin.HandlerFunc
}

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandlers(userUsecase domain.UserUsecase) UserHandlers {
	return &UserController{UserUsecase: userUsecase}
}

func (ctrl *UserController) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users := ctrl.UserUsecase.FindAll()
		c.JSON(200, users)
	}
}

// func (ctrl *UserController) DeleteAllUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		err := ctrl.UserUsecase.DeleteAll()
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "no users found"})
// 			return
// 		}
// 		c.JSON(http.StatusNoContent, gin.H{"message": "all users deleted"})
// 	}
// }

func (ctrl *UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		user, exist := ctrl.UserUsecase.FindByUsername(username)

		if !exist {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(200, domain.UserRes{ID: user.ID, Username: user.Username, Role: user.Role})
	}
}

func (ctrl *UserController) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		user, exists := ctrl.UserUsecase.FindByUsername(username)

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.BindJSON(&user)
		newUser, err := ctrl.UserUsecase.Update(username, user)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, newUser)
	}
}

func (ctrl *UserController) PromoteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		newUser, err := ctrl.UserUsecase.PromoteUser(username)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, newUser)
	}
}

func (ctrl *UserController) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		userData, _ := c.Get("user")

		if username != userData.(*domain.User).Username {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		err := ctrl.UserUsecase.Delete(username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{"message": "user deleted"})
	}
}
