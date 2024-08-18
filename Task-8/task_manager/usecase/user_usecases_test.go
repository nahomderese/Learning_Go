package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/infrastructure"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseSuite struct {
	suite.Suite
	mockRepo      *mocks.UserRepository
	userUsecase   domain.UserUsecase
	mockUser      domain.User
	mockUserRes   domain.UserRes
	mockTask      domain.Task
	mockEmptyTask domain.Task
	mockUserID    primitive.ObjectID
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = mocks.NewUserRepository(suite.T())
	suite.mockUserID = primitive.NewObjectID()
	suite.mockUser = domain.User{
		ID:       suite.mockUserID,
		Username: "username",
		Role:     "admin",
		Password: "passwordpass",
	}
	suite.mockUserRes = domain.UserRes{
		ID:       suite.mockUserID,
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

	suite.userUsecase = NewUserUseCase(suite.mockRepo)

	pass, err := infrastructure.HashPassword(suite.mockUser.Password)

	if err != nil {
		suite.T().Fatal(err)
	}

	suite.mockUser.Password = pass

	suite.mockEmptyTask = domain.Task{}
}

// func (suite *UserUsecaseSuite) TestSuccessPromoteUser() {

// 	suite.mockRepo.On("FindByUsername", mock.Anything).Return(suite.mockUser, true)
// 	suite.mockRepo.On("Update", mock.Anything, mock.Anything).Return(suite.mockUser, nil)

// 	_, err := suite.userUsecase.PromoteUser(suite.mockUser.Username)

// 	suite.Error(err)

// 	suite.mockRepo.AssertExpectations(suite.T())

// }

func (suite *UserUsecaseSuite) TestSuccessDelete() {

	suite.mockRepo.On("Delete", suite.mockUser.Username).Return(nil).Once()

	err := suite.userUsecase.Delete(suite.mockUser.Username)

	suite.NoError(err)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessFindAll() {

	suite.mockRepo.On("FindAll").Return([]domain.User{suite.mockUser}).Once()

	tasks := suite.userUsecase.FindAll()

	suite.IsType([]domain.UserRes{}, tasks)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessFindByID() {

	suite.mockRepo.On("FindUser", mock.Anything).Return(suite.mockUser, nil)

	task, err := suite.userUsecase.FindUser(suite.mockTask.ID.Hex())

	suite.NoError(err)

	suite.Equal(suite.mockUserRes, task)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessFindByUsername() {

	suite.mockRepo.On("FindByUsername", suite.mockUser.Username).Return(suite.mockUser, true)

	user, exists := suite.userUsecase.FindByUsername(suite.mockUser.Username)

	suite.True(exists)
	suite.Equal(suite.mockUserRes, user)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessCreate() {

	suite.mockRepo.On("Save", suite.mockUser).Return(suite.mockUser, nil)

	task, err := suite.userUsecase.CreateUser(suite.mockUser)

	suite.NoError(err)

	suite.Equal(suite.mockUserRes, task)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessUpdate() {

	suite.mockRepo.On("Update", suite.mockUser.Username, mock.Anything).Return(suite.mockUser, nil).Once()

	user, err := suite.userUsecase.Update(suite.mockUser.Username, suite.mockUserRes)

	suite.NoError(err)

	suite.Equal(suite.mockUserRes, user)

	suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *UserUsecaseSuite) TestSuccessLogin() {

	fmt.Println(suite.mockUser.Password)

	suite.mockRepo.On("FindByUsername", suite.mockUser.Username).Return(suite.mockUser, true)

	token, err := suite.userUsecase.Login(suite.mockUser.Username, "passwordpass")

	suite.NoError(err)

	suite.NotEmpty(token)

	suite.mockRepo.AssertExpectations(suite.T())

}

func TestUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
