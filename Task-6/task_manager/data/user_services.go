package data

import (
	"context"
	"errors"
	"log"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskRepository is an interface for managing tasks.

type UserRepository interface {
	Save(user *models.User) (models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindAll() []models.User
	Delete(username string) error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

// Delete implements UserRepository.
func (repo *MongoUserRepository) Delete(username string) error {

	_, err := repo.collection.DeleteOne(context.TODO(), bson.M{"username": username})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("user not found")
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

// FindAll implements UserRepository.
func (repo *MongoUserRepository) FindAll() []models.User {

	cursor, err := repo.collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var elem models.User
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, elem)
	}

	return users
}

// FindByID implements UserRepository.
func (repo *MongoUserRepository) FindByUsername(username string) (*models.User, error) {

	var user models.User
	err := repo.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		} else {
			log.Fatal(err)
		}
	}

	return &user, nil
}

// Save implements UserRepository.
func (repo *MongoUserRepository) Save(user *models.User) (models.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := repo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

// Constructor functions

func NewMongoUserRepository(collection *mongo.Collection) UserRepository {
	return &MongoUserRepository{collection}
}
