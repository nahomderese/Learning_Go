package controllers

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthSuite struct {
	suite.Suite
	mockUsecase *mocks.UserUsecase
	ctrl        AuthHandlers
	mockUser    domain.User
	mockUserRes domain.UserRes
	mockUserID  primitive.ObjectID
	router      *gin.Engine
	token       string
}

func (suite *AuthSuite) SetupTest() {
	suite.mockUsecase = mocks.NewUserUsecase(suite.T())
	suite.mockUserID = primitive.NewObjectID()
	suite.mockUser = domain.User{
		ID:       suite.mockUserID,
		Username: "username",
		Role:     "admin",
		Password: "password",
	}
	suite.mockUserRes = domain.UserRes{
		ID:       suite.mockUserID,
		Username: "username",
		Role:     "admin",
	}
	suite.token = "loremipsumidorLoremipsumIdor"

	// controller
	suite.ctrl = NewAuthHandlers(suite.mockUsecase)

	// gin router
	suite.router = gin.Default()
	authRoutes := suite.router.Group("auth")
	authRoutes.POST("/register", suite.ctrl.SignUp())
	authRoutes.POST("/login", suite.ctrl.Login())

}

func (suite *AuthSuite) TestSignUpSuccess() {

	suite.mockUsecase.On("Signup", mock.Anything, mock.Anything).Return(suite.mockUserRes, nil)

	req, _ := http.NewRequest("POST", "/auth/register", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"username":"username","password":"password"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

}
func (suite *AuthSuite) TestSignUpAlreadyExist() {

	suite.mockUsecase.On("Signup", mock.Anything, mock.Anything).Return(domain.UserRes{}, errors.New("Username already exists"))

	req, _ := http.NewRequest("POST", "/auth/register", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"username":"username","password":"password"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusConflict, w.Code)

}
func (suite *AuthSuite) TestSignUpError() {

	suite.mockUsecase.On("Signup", mock.Anything, mock.Anything).Return(domain.UserRes{}, errors.New("Oops! Something went wrong"))

	req, _ := http.NewRequest("POST", "/auth/register", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"username":"username","password":"password"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)

}

func (suite *AuthSuite) TestLoginSuccess() {

	suite.mockUsecase.On("Login", mock.Anything, mock.Anything).Return(suite.token, nil)

	req, _ := http.NewRequest("POST", "/auth/login", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"username":"username","password":"password"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}
func (suite *AuthSuite) TestLoginError() {

	suite.mockUsecase.On("Login", mock.Anything, mock.Anything).Return("", errors.New("Invalid username or password"))

	req, _ := http.NewRequest("POST", "/auth/login", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"username":"username","password":"password"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusUnauthorized, w.Code)

}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}
