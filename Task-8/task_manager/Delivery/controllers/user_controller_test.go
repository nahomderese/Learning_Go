package controllers

import (
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerSuite struct {
	suite.Suite
	mockCollection *mocks.MongoCollection
	mockUser       domain.User
	mockUserID     primitive.ObjectID
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
}
