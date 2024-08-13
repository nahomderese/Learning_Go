package controllers

import (
	"net/http"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlers interface {
	SignUp() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type AuthController struct {
	UserUsecase domain.UserUsecase
}

func (ctrl *AuthController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		err := c.ShouldBindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		// if no users in the db, create an admin user
		users := ctrl.UserUsecase.FindAll()

		role := "regular"
		if len(users) == 0 {
			role = "admin"
		}

		// check if user already exists
		_, exists := ctrl.UserUsecase.FindByUsername(user.Username)

		if exists {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}

		userData := domain.User{Username: user.Username, Password: string(hash), Role: role}
		newUser, error := ctrl.UserUsecase.CreateUser(userData)

		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}

		c.JSON(http.StatusOK, newUser)

	}
}

func (ctrl *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		err := c.BindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := ctrl.UserUsecase.Login(user.Username, user.Password)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})

	}
}
