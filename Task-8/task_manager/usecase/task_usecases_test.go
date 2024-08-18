package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseSuite struct {
	suite.Suite
	mockRepo      *mocks.TaskRepository
	mockUser      domain.User
	mockTask      domain.Task
	mockEmptyTask domain.Task
	mockUserID    primitive.ObjectID
}

func (suite *TaskUsecaseSuite) SetupTest() {
	suite.mockRepo = mocks.NewTaskRepository(suite.T())
	suite.mockUser = domain.User{
		ID:       primitive.NewObjectID(),
		Username: "username",
		Role:     "admin",
		Password: "password",
	}
	suite.mockUserID = primitive.NewObjectID()
	suite.mockTask = domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "title",
		Description: "description",
		DueDate:     time.Now(),
		Status:      "status",
		UserId:      suite.mockUserID.Hex(),
	}

	suite.mockEmptyTask = domain.Task{}
}

func (suite *TaskUsecaseSuite) TestSuccessCreate() {

	suite.mockRepo.On("Save", mock.Anything, mock.Anything).Return(suite.mockTask, nil).Once()

	tu := NewTaskUseCase(suite.mockRepo)

	task, err := tu.Save(context.TODO(), suite.mockTask)

	suite.NoError(err)

	suite.Equal(suite.mockTask, task)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *TaskUsecaseSuite) TestSuccessGetAll() {
	suite.mockRepo.On("FindAll", mock.Anything, mock.Anything).Return([]domain.Task{suite.mockTask}, nil).Once()

	tu := NewTaskUseCase(suite.mockRepo)

	tasks, err := tu.FindAll(context.Background(), suite.mockUser)

	suite.NoError(err)

	suite.Equal([]domain.Task{suite.mockTask}, tasks)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestSuccessGetByID() {
	suite.mockRepo.On("FindByID", mock.Anything, mock.Anything).Return(suite.mockTask, nil).Once()

	tu := NewTaskUseCase(suite.mockRepo)

	task, err := tu.FindByID(context.Background(), suite.mockTask.ID)

	suite.NoError(err)

	suite.Equal(suite.mockTask, task)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestSuccessDelete() {
	suite.mockRepo.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()

	tu := NewTaskUseCase(suite.mockRepo)

	err := tu.Delete(context.Background(), suite.mockTask.ID)

	suite.NoError(err)

	suite.mockRepo.AssertExpectations(suite.T())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}
