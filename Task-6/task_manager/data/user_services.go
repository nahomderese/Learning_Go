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
	Save(user models.User) (models.UserRes, error)
	FindById(id primitive.ObjectID) (models.UserRes, error)
	FindUser(id string) (models.User, error)
	FindAll() []models.UserRes
	Delete(username string) error
	DeleteAll() error
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

// DeleteAll implements UserRepository.
func (repo *MongoUserRepository) DeleteAll() error {

	_, err := repo.collection.DeleteMany(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
		return errors.New("error deleting all users")
	}

	return nil
}

// FindAll implements UserRepository.
func (repo *MongoUserRepository) FindAll() []models.UserRes {

	cursor, err := repo.collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var users []models.UserRes

	for cursor.Next(context.Background()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, models.UserRes{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		})
	}

	return users

}

// FindByUsername implements UserRepository.
func (repo *MongoUserRepository) FindByUsername(username string) (models.UserRes, error) {

	var user models.User
	err := repo.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)

	return handleUserError(user, err)
}

// FindByID implements UserRepository.
func (repo *MongoUserRepository) FindById(id primitive.ObjectID) (models.UserRes, error) {

	var user models.User
	err := repo.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	return handleUserError(user, err)
}

// FindUser implements UserRepository.
func (repo *MongoUserRepository) FindUser(id string) (models.User, error) {

	var user models.User

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.User{}, errors.New("invalid id")
	}

	repo.collection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return models.User{}, errors.New("user not found")
	} else {
		log.Fatal(err)
	}

	return user, nil
}

// Save implements UserRepository.
func (repo *MongoUserRepository) Save(user models.User) (models.UserRes, error) {
	user.ID = primitive.NewObjectID()
	_, err := repo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return models.UserRes{}, err
	}
	return models.UserRes{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}, nil

}

// Constructor functions

func NewMongoUserRepository(collection *mongo.Collection) UserRepository {
	return &MongoUserRepository{collection}
}

func handleUserError(user models.User, err error) (models.UserRes, error) {

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.UserRes{}, errors.New("user not found")
		} else {
			log.Fatal(err)
		}
	}

	return models.UserRes{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}, nil

}
