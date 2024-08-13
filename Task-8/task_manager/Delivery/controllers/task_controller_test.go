package controllers

import (
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/infrastructure"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerSuite struct {
	suite.Suite
	mockUserUsecase *mocks.UserUsecase
	mockTaskUsecase *mocks.TaskUsecase
	ctrl            TaskHandlers
	mockUser        domain.User
	mockUserRes     domain.UserRes
	mockUserID      primitive.ObjectID
	router          *gin.Engine
	token           string
}

func (suite *TaskControllerSuite) SetupTest() {
	suite.mockUserUsecase = mocks.NewUserUsecase(suite.T())
	suite.mockTaskUsecase = mocks.NewTaskUsecase(suite.T())
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
	suite.ctrl = NewTaskHandlers(suite.mockTaskUsecase, suite.mockUserUsecase)

	// gin router
	suite.router = gin.Default()
	tasksRoutes := suite.router.Group("tasks")
	tasksRoutes.GET("", suite.ctrl.GetAllTasks())
	tasksRoutes.GET(":id", suite.ctrl.GetTaskById())

	tasksRoutes.Use(infrastructure.AdminAuthMiddleware())
	tasksRoutes.POST("", suite.ctrl.CreateTask())
	tasksRoutes.PUT(":id", suite.ctrl.UpdateTask())
	tasksRoutes.DELETE(":id", suite.ctrl.DeleteTask())

}
