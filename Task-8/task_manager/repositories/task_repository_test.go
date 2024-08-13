package repositories

import (
	"errors"
	"testing"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreate(t *testing.T) {

	collectionHelper := mocks.NewMongoCollection(t)

	mockUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "username",
		Role:     "admin",
		Password: "password",
	}

	mockEmptyUser := domain.User{}
	mockUserID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {

		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("domain.User")).Return(&mongo.InsertOneResult{
			InsertedID: mockUserID,
		}, nil).Once()

		ur := NewUserRepository(collectionHelper)

		usr, err := ur.Save(mockUser)

		assert.NoError(t, err)

		assert.Equal(t, mockUserID, usr.ID)

		collectionHelper.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("domain.User")).Return(&mongo.InsertOneResult{
			InsertedID: mockUserID,
		}, errors.New("error")).Once()

		ur := NewUserRepository(collectionHelper)

		usr, err := ur.Save(mockEmptyUser)

		assert.Error(t, err)

		assert.Equal(t, mockEmptyUser, usr)

		collectionHelper.AssertExpectations(t)
	})

}
