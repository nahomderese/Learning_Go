package controllers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/infrastructure"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerSuite struct {
	suite.Suite
	mockUserUsecase *mocks.UserUsecase
	mockUserRepo    *mocks.UserRepository
	mockTaskUsecase *mocks.TaskUsecase
	ctrl            TaskHandlers
	mockUser        domain.User
	mockTask        domain.Task
	mockUserRes     domain.UserRes
	mockUserID      primitive.ObjectID
	mockTaskID      primitive.ObjectID
	router          *gin.Engine
	token           string
}

func (suite *TaskControllerSuite) SetupTest() {
	suite.mockUserUsecase = mocks.NewUserUsecase(suite.T())
	suite.mockTaskUsecase = mocks.NewTaskUsecase(suite.T())
	suite.mockUserRepo = mocks.NewUserRepository(suite.T())
	suite.mockUserID = primitive.NewObjectID()
	suite.mockTaskID = primitive.NewObjectID()
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
	suite.token = "loremipsumidorLoremipsum"

	suite.mockTask = domain.Task{
		ID:          suite.mockTaskID,
		Title:       "title",
		Description: "description",
		DueDate:     time.Now(),
		Status:      "pending",
		UserId:      suite.mockUserID.Hex(),
	}

	token, err := infrastructure.GenerateToken(suite.mockUser.Username, suite.mockUser.ID.Hex())
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.token = token
	suite.token = "Bearer " + suite.token

	// controller
	suite.ctrl = NewTaskHandlers(suite.mockTaskUsecase, suite.mockUserUsecase)

	// gin router
	suite.router = gin.Default()
	tasksRoutes := suite.router.Group("tasks")
	tasksRoutes.Use(infrastructure.AuthMiddleware(suite.mockUserRepo))
	tasksRoutes.GET("", suite.ctrl.GetAllTasks())
	tasksRoutes.GET(":id", suite.ctrl.GetTaskById())

	tasksRoutes.Use(infrastructure.AdminAuthMiddleware())
	tasksRoutes.POST("", suite.ctrl.CreateTask())
	tasksRoutes.PUT(":id", suite.ctrl.UpdateTask())
	tasksRoutes.DELETE(":id", suite.ctrl.DeleteTask())

	suite.mockUserRepo.On("FindByUsername", mock.Anything).Return(suite.mockUser, true)

}

func (suite *TaskControllerSuite) TestGetAllTasksSuccess() {
	suite.mockTaskUsecase.On("FindAll", mock.Anything, mock.Anything).Return([]domain.Task{suite.mockTask}, nil)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), suite.mockTask.Title)
}

func (suite *TaskControllerSuite) TestGetTaskByIdSuccess() {
	suite.mockTaskUsecase.On("FindByID", mock.Anything, mock.Anything).Return(suite.mockTask, nil)

	req, _ := http.NewRequest("GET", "/tasks/"+suite.mockTaskID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), suite.mockTask.Title)
}

func (suite *TaskControllerSuite) TestCreateTaskSuccess() {
	suite.mockTaskUsecase.On("Save", mock.Anything, mock.Anything).Return(suite.mockTask, nil)
	suite.mockUserUsecase.On("FindUser", mock.Anything).Return(suite.mockUserRes, nil)
	fmt.Println(suite.token)
	req, _ := http.NewRequest("POST", "/tasks", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	req.Body = io.NopCloser(strings.NewReader(`{"title":"title","description":"description","due_date":"2006-01-02T15:04:05Z","status":"pending"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *TaskControllerSuite) TestUpdateTaskSuccess() {
	suite.mockTaskUsecase.On("FindByID", mock.Anything, mock.Anything).Return(suite.mockTask, nil)
	suite.mockTaskUsecase.On("Save", mock.Anything, mock.Anything).Return(suite.mockTask, nil)

	req, _ := http.NewRequest("PUT", "/tasks/"+suite.mockTaskID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	req.Body = io.NopCloser(strings.NewReader(`{"title":"title","description":"description","due_date":"2006-01-02T15:04:05Z","status":"pending"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *TaskControllerSuite) TestDeleteTaskSuccess() {
	suite.mockTaskUsecase.On("Delete", mock.Anything, mock.Anything).Return(nil)

	req, _ := http.NewRequest("DELETE", "/tasks/"+suite.mockTaskID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", suite.token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNoContent, w.Code)
}

func TestTaskSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
