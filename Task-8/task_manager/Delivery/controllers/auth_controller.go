package controllers

import (
	"net/http"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/gin-gonic/gin"
)

type AuthHandlers interface {
	SignUp() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type AuthController struct {
	UserUsecase domain.UserUsecase
}

func NewAuthHandlers(userUsecase domain.UserUsecase) AuthHandlers {
	return &AuthController{UserUsecase: userUsecase}
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

		newUser, e := ctrl.UserUsecase.Signup(user.Username, user.Password)

		if e != nil {
			if e.Error() == "Username already exists" {
				c.JSON(http.StatusConflict, gin.H{"error": e.Error()})
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
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
