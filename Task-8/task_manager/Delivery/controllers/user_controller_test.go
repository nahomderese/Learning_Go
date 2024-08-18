package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/infrastructure"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerSuite struct {
	suite.Suite
	mockCollection  *mocks.MongoCollection
	mockUser        domain.User
	mockUserID      primitive.ObjectID
	mockUserUsecase *mocks.UserUsecase
	mockUserRepo    *mocks.UserRepository
	ctrl            UserHandlers
	mockUserRes     domain.UserRes
	router          *gin.Engine
	token           string
}

func (suite *UserControllerSuite) SetupTest() {
	suite.mockCollection = mocks.NewMongoCollection(suite.T())
	suite.mockUser = domain.User{
		ID:       primitive.NewObjectID(),
		Username: "username",
		Role:     "admin",
		Password: "password",
	}
	suite.mockUserID = primitive.NewObjectID()
	suite.mockUserUsecase = mocks.NewUserUsecase(suite.T())
	suite.mockUserRepo = mocks.NewUserRepository(suite.T())
	suite.mockUserRes = domain.UserRes{
		ID:       suite.mockUserID,
		Username: "username",
		Role:     "admin",
	}

	token, err := infrastructure.GenerateToken(suite.mockUser.Username, suite.mockUser.ID.Hex())
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.token = token
	suite.token = "Bearer " + suite.token

	suite.ctrl = NewUserHandlers(suite.mockUserUsecase)

	suite.router = gin.Default()
	usersRoutes := suite.router.Group("users")
	usersRoutes.Use(infrastructure.AuthMiddleware(suite.mockUserRepo))
	usersRoutes.PUT(":username", suite.ctrl.UpdateUser())
	usersRoutes.GET(":username", suite.ctrl.GetUser())
	usersRoutes.DELETE(":username", suite.ctrl.DeleteUser())

	usersRoutes.Use(infrastructure.AdminAuthMiddleware())
	usersRoutes.GET("", suite.ctrl.GetAllUsers())
	usersRoutes.PATCH("/promote/:username", suite.ctrl.PromoteUser())

	suite.mockUserRepo.On("FindByUsername", mock.Anything).Return(suite.mockUser, true)

}

func (suite *UserControllerSuite) TestGetAllUsers() {
	suite.mockUserUsecase.On("FindAll").Return([]domain.UserRes{suite.mockUserRes})

	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

}

func (suite *UserControllerSuite) TestGetUser() {
	suite.mockUserUsecase.On("FindByUsername", suite.mockUser.Username).Return(suite.mockUserRes, true)

	req, _ := http.NewRequest("GET", "/users/username", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

}

func (suite *UserControllerSuite) TestUpdateUser() {
	suite.mockUserUsecase.On("FindByUsername", mock.Anything).Return(suite.mockUserRes, true)
	suite.mockUserUsecase.On("Update", mock.Anything, mock.Anything).Return(suite.mockUserRes, nil)

	req, _ := http.NewRequest("PUT", "/users/username", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	req.Body = io.NopCloser(strings.NewReader(`{"username":"username","password":"password"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusAccepted, w.Code)

}

func (suite *UserControllerSuite) TestDeleteUser() {
	suite.mockUserUsecase.On("Delete", mock.Anything).Return(nil)

	req, _ := http.NewRequest("DELETE", "/users/username", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNoContent, w.Code)
}

func (suite *UserControllerSuite) TestPromoteUser() {
	suite.mockUserUsecase.On("PromoteUser", mock.Anything).Return(suite.mockUserRes, nil)

	req, _ := http.NewRequest("PATCH", "/users/promote/username", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusAccepted, w.Code)

}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
