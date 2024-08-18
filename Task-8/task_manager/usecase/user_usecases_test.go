package usecase

import (
	"time"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseSuite struct {
	suite.Suite
	mockRepo      *mocks.UserRepository
	mockUser      domain.User
	mockUserRes   domain.UserRes
	mockTask      domain.Task
	mockEmptyTask domain.Task
	mockUserID    primitive.ObjectID
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = mocks.NewUserRepository(suite.T())
	suite.mockUser = domain.User{
		ID:       primitive.NewObjectID(),
		Username: "username",
		Role:     "admin",
		Password: "password",
	}
	suite.mockUserRes = domain.UserRes{
		ID:       primitive.NewObjectID(),
		Username: "username",
		Role:     "admin",
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

func (suite *UserUsecaseSuite) TestSuccessPromoteUser() {

	suite.mockRepo.On("FindByUsername", suite.mockUser.Username).Return(suite.mockUser, true).Once()
	suite.mockRepo.On("Update", suite.mockUser.Username, suite.mockUser).Return(suite.mockUser, nil).Once()

	tu := NewUserUseCase(suite.mockRepo)

	user, err := tu.PromoteUser(suite.mockUser.Username)

	suite.NoError(err)

	suite.Equal(suite.mockUser, user)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessDelete() {

	suite.mockRepo.On("Delete", suite.mockUser.Username).Return(nil).Once()

	tu := NewUserUseCase(suite.mockRepo)

	err := tu.Delete(suite.mockUser.Username)

	suite.NoError(err)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessDeleteAll() {

	suite.mockRepo.On("DeleteAll").Return(nil).Once()

	tu := NewUserUseCase(suite.mockRepo)

	err := tu.DeleteAll()

	suite.NoError(err)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessFindAll() {

	suite.mockRepo.On("FindAll").Return([]domain.Task{suite.mockTask}).Once()

	tu := NewUserUseCase(suite.mockRepo)

	tasks := tu.FindAll()

	suite.Equal([]domain.Task{suite.mockTask}, tasks)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessFindByID() {

	suite.mockRepo.On("FindByID", suite.mockTask.ID).Return(suite.mockTask, nil).Once()

	tu := NewUserUseCase(suite.mockRepo)

	task, err := tu.FindUser(suite.mockTask.ID.Hex())

	suite.NoError(err)

	suite.Equal(suite.mockTask, task)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessFindByUsername() {

	suite.mockRepo.On("FindByUsername", suite.mockUser.Username).Return(suite.mockUser, true).Once()

	tu := NewUserUseCase(suite.mockRepo)

	user, exists := tu.FindByUsername(suite.mockUser.Username)

	suite.True(exists)

	suite.Equal(suite.mockUser, user)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessCreate() {

	suite.mockRepo.On("Save", suite.mockTask).Return(suite.mockTask, nil).Once()

	tu := NewUserUseCase(suite.mockRepo)

	task, err := tu.CreateUser(suite.mockUser)

	suite.NoError(err)

	suite.Equal(suite.mockTask, task)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessUpdate() {

	suite.mockRepo.On("Update", suite.mockUser.Username, suite.mockUser).Return(suite.mockUser, nil).Once()

	tu := NewUserUseCase(suite.mockRepo)

	user, err := tu.Update(suite.mockUser.Username, suite.mockUserRes)

	suite.NoError(err)

	suite.Equal(suite.mockUser, user)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessLogin() {

	suite.mockRepo.On("FindByUsername", suite.mockUser.Username).Return(suite.mockUser, true).Once()

	tu := NewUserUseCase(suite.mockRepo)

	user, err := tu.Login(suite.mockUser.Username, suite.mockUser.Password)

	suite.NoError(err)

	suite.Equal(suite.mockUser, user)

	suite.mockRepo.AssertExpectations(suite.T())

}
