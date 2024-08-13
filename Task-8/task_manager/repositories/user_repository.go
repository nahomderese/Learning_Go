package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskRepository is an interface for managing tasks.

type MongoUserRepository struct {
	collection infrastructure.MongoCollection
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
func (repo *MongoUserRepository) FindAll() []domain.User {

	cursor, err := repo.collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var users []domain.User = make([]domain.User, 0)

	for cursor.Next(context.Background()) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users

}

// FindByUsername implements UserRepository.
func (repo *MongoUserRepository) FindByUsername(username string) (domain.User, bool) {

	var user domain.User

	err := repo.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)

	if err != nil {
		return domain.User{}, false
	}

	return user, user.Username == username

}

// FindByID implements UserRepository.
func (repo *MongoUserRepository) FindById(id primitive.ObjectID) (domain.User, error) {

	var user domain.User
	err := repo.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	return handleUserError(user, err)
}

// FindUser implements UserRepository.
func (repo *MongoUserRepository) FindUser(id string) (domain.User, error) {

	var user domain.User

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return domain.User{}, errors.New("invalid id")
	}

	repo.collection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&user)

	if user.ID.Hex() != id {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil

}

// Save implements UserRepository.
func (repo *MongoUserRepository) Save(user domain.User) (domain.User, error) {
	res, err := repo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       res.InsertedID.(primitive.ObjectID),
		Username: user.Username,
		Role:     user.Role,
	}, nil

}

// Save implements UserRepository.
func (repo *MongoUserRepository) Update(username string, user domain.User) (domain.User, error) {
	_, err := repo.collection.UpdateOne(context.TODO(), bson.M{"username": username}, bson.M{"$set": bson.M{"role": user.Role}})

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}, nil

}

// Constructor functions

func NewUserRepository(collection infrastructure.MongoCollection) domain.UserRepository {
	return &MongoUserRepository{collection}
}

func handleUserError(user domain.User, err error) (domain.User, error) {

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		} else {
			log.Fatal(err)
		}
	}

	return domain.User{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}, nil

}
