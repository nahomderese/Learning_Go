package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlers interface {
	GetAllUsers() gin.HandlerFunc
	DeleteAllUsers() gin.HandlerFunc
	GetUserByUsername() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	PromoteUser() gin.HandlerFunc

	SignUp() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type UserController struct {
	UserRepo data.UserRepository
}

func (ctrl *UserController) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users := ctrl.UserRepo.FindAll()
		c.JSON(200, users)
	}
}

func (ctrl *UserController) DeleteAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := ctrl.UserRepo.DeleteAll()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "users not found"})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "users deleted"})
	}
}

func (ctrl *UserController) GetUserByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		user, err := ctrl.UserRepo.FindByUsername(username)

		userData, _ := c.Get("user")
		if username != userData.(*models.User).Username {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

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

		userData, _ := c.Get("user")

		if username != userData.(*models.User).Username {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		user, err := ctrl.UserRepo.FindUser(username)
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

func (ctrl *UserController) PromoteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		user, err := ctrl.UserRepo.FindUser(username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.BindJSON(&user)

		if user.Role == "admin" {
			c.JSON(http.StatusConflict, gin.H{"error": "user is already an admin"})
			return
		}

		user.Role = "admin"
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

		userData, _ := c.Get("user")

		if username != userData.(*models.User).Username {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		err := ctrl.UserRepo.Delete(username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(200, gin.H{"message": "user deleted"})
	}
}

func (ctrl *UserController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

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

		// if no users in the db, create an admin user
		users := ctrl.UserRepo.FindAll()

		role := "regular"
		if len(users) == 0 {
			role = "admin"
		}

		newUser := models.User{Username: user.Username, Password: string(hash), Role: role}
		ctrl.UserRepo.Save(newUser)

		user.Role = role

		c.JSON(http.StatusOK, user)

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
		userFromDB, err := ctrl.UserRepo.FindUser(user.Username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
			return
		}

		// generate token
		token, err := GenerateToken(userFromDB.Username, userFromDB.Role)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})

	}
}

func GenerateToken(username, role string) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().AddDate(0, 0, 1).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
