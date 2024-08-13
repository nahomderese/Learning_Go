package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskSuite struct {
	suite.Suite
	mockCollection *mocks.MongoCollection
	mockUser       domain.User
	mockTask       domain.Task
	mockEmptyTask  domain.Task
	mockUserID     primitive.ObjectID
}

func (suite *TaskSuite) SetupTest() {
	suite.mockCollection = mocks.NewMongoCollection(suite.T())
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

func (suite *TaskSuite) TestSuccessCreate() {
	suite.mockCollection.On("InsertOne", mock.Anything, mock.AnythingOfType("domain.Task")).Return(&mongo.InsertOneResult{
		InsertedID: suite.mockUserID,
	}, errors.New("error")).Once()

	tr := NewTaskRepository(suite.mockCollection)

	task, err := tr.Save(context.TODO(), suite.mockTask)

	suite.NoError(err)

	suite.Equal(suite.mockTask.ID, task.ID)

	suite.mockCollection.AssertExpectations(suite.T())
}

func (suite *TaskSuite) TestErrorCreate() {
	suite.mockCollection.On("InsertOne", mock.Anything, mock.AnythingOfType("domain.Task")).Return(&mongo.InsertOneResult{
		InsertedID: suite.mockUserID,
	}, errors.New("error")).Once()

	tr := NewTaskRepository(suite.mockCollection)

	_, err := tr.Save(context.TODO(), suite.mockTask)

	suite.Error(err)

	suite.mockCollection.AssertExpectations(suite.T())
}

func (suite *TaskSuite) TestSuccessGetByID() {
	suite.mockCollection.On("FindOne", mock.Anything, mock.AnythingOfType("primitive.M")).Return(&mongo.SingleResult{}, nil).Once()

	tr := NewTaskRepository(suite.mockCollection)

	_, err := tr.FindByID(context.TODO(), suite.mockTask.ID)

	suite.NoError(err)

	suite.mockCollection.AssertExpectations(suite.T())
}

func (suite *TaskSuite) TestErrorGetByID() {
	suite.mockCollection.On("FindOne", mock.Anything, mock.AnythingOfType("primitive.M")).Return(&mongo.SingleResult{}, errors.New("error")).Once()

	tr := NewTaskRepository(suite.mockCollection)

	_, err := tr.FindByID(context.TODO(), suite.mockTask.ID)

	suite.Error(err)

	suite.mockCollection.AssertExpectations(suite.T())
}

func (suite *TaskSuite) TestSuccessGetAll() {
	suite.mockCollection.On("Find", mock.Anything, mock.AnythingOfType("primitive.M")).Return(&mongo.Cursor{}, nil).Once()

	tr := NewTaskRepository(suite.mockCollection)

	tasks, err := tr.FindAll(context.TODO(), suite.mockUser)

	suite.NoError(err)

	suite.Equal([]domain.Task{}, tasks)

	suite.mockCollection.AssertExpectations(suite.T())
}

func (suite *TaskSuite) TestErrorGetAll() {
	suite.mockCollection.On("Find", mock.Anything, mock.AnythingOfType("primitive.M")).Return(&mongo.Cursor{}, errors.New("error")).Once()

	tr := NewTaskRepository(suite.mockCollection)

	tasks, err := tr.FindAll(context.TODO(), suite.mockUser)

	suite.Error(err)

	suite.Equal([]domain.Task{}, tasks)

	suite.mockCollection.AssertExpectations(suite.T())
}
