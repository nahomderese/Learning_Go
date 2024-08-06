package controllers

import (
	"net/http"
	"os"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlers interface {
	GetAllUsers() gin.HandlerFunc
	GetUserByUsername() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	CreateUser() gin.HandlerFunc

	Login() gin.HandlerFunc
}

type UserController struct {
	UserRepo data.UserRepository
}

func (ctrl *UserController) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks := ctrl.UserRepo.FindAll()
		c.JSON(200, tasks)
	}
}

func (ctrl *UserController) GetUserByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		user, err := ctrl.UserRepo.FindByUsername(username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func (ctrl *UserController) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		user, err := ctrl.UserRepo.FindByUsername(username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.BindJSON(&user)
		newUser, err := ctrl.UserRepo.Save(user)

		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newUser)
	}
}

func (ctrl *UserController) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		err := ctrl.UserRepo.Delete(username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(200, gin.H{"message": "user deleted"})
	}
}

func (ctrl *UserController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}

		err := c.BindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		newUser := models.User{Username: user.Username, Password: string(hash), Role: user.Role}

		ctrl.UserRepo.Save(&newUser)
		c.JSON(200, user)

	}
}

func (ctrl *UserController) Login() gin.HandlerFunc {
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

		// get user from db
		userFromDB, err := ctrl.UserRepo.FindByUsername(user.Username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		}

		// generate token
		token, err := GenerateToken(userFromDB.Username, userFromDB.Role)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"token": token})

	}
}

func GenerateToken(username, role string) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
