package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlers interface {
	GetAllUsers() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	DeleteAllUser() gin.HandlerFunc
	PromoteUser() gin.HandlerFunc
}

type AuthHandlers interface {
	SignUp() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type UserController struct {
	UserRepo data.UserRepository
}

type AuthController struct {
	UserRepo data.UserRepository
}

func (ctrl *UserController) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users := ctrl.UserRepo.FindAll()
		c.JSON(200, users)
	}
}

func (ctrl *UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := ctrl.UserRepo.FindUser(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(200, models.UserRes{ID: users.ID, Username: users.Username, Role: users.Role})
	}
}

func (ctrl *UserController) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id_ := c.Param("username")

		id, err := primitive.ObjectIDFromHex(id_)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		user, err := ctrl.UserRepo.FindById(id)

		userData, _ := c.Get("user")
		if id_ != userData.(*models.User).ID.Hex() && userData.(*models.User).Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(200, models.UserRes{ID: user.ID, Username: user.Username, Role: user.Role})
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

		user, exist := ctrl.UserRepo.FindByUsername(username)
		if !exist {
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

		user, err := ctrl.UserRepo.FindByUsername(username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		if user.Role == "admin" {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "user is already an admin"})
			return
		}

		user.Role = "admin"
		newUser, err := ctrl.UserRepo.Update(user)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
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

func (ctrl *AuthController) SignUp() gin.HandlerFunc {
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


		userData := models.User{Username: user.Username, Password: string(hash), Role: role}
		newUser, error := ctrl.UserRepo.Save(userData)

    // check if user already exists
		_, err = ctrl.UserRepo.FindUser(user.ID.Hex())
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
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

		// get user from db
		userFromDB, exist := ctrl.UserRepo.FindByUsername(user.Username)
		if !exist {
			c.JSON(http.StatusNotFound, gin.H{"error": "username or password is incorrect"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect"})
			return
		}

		// generate token
		token, err := GenerateToken(userFromDB.Username, userFromDB.ID.Hex())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})

	}
}

func GenerateToken(username, id string) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"exp":      time.Now().AddDate(0, 0, 1).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
