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

type TaskRepository interface {
	Save(c context.Context, task models.Task) (models.Task, error)
	FindByID(c context.Context, id primitive.ObjectID) (models.Task, error)
	FindAll(c context.Context, user models.User) []models.Task
	Delete(c context.Context, id primitive.ObjectID) error
}

type MongoTaskRepository struct {
	collection *mongo.Collection
}

// Delete implements TaskRepository.
func (repo *MongoTaskRepository) Delete(c context.Context, id primitive.ObjectID) error {

	result, _ := repo.collection.DeleteOne(c, bson.M{"_id": id})

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

// FindAll implements TaskRepository.
func (repo *MongoTaskRepository) FindAll(c context.Context, user models.User) []models.Task {

	query := bson.D{}

	if user.Role != "admin" {
		query = bson.D{{Key: "user_id", Value: user.ID.Hex()}}
	}

	cursor, err := repo.collection.Find(c, query)

	if err != nil {
		log.Fatal(err)
	}

	var tasks []models.Task = make([]models.Task, 0)
	if err = cursor.All(c, &tasks); err != nil {
		log.Fatal(err)
	}

	return tasks
}

// FindByID implements TaskRepository.
func (repo *MongoTaskRepository) FindByID(c context.Context, id primitive.ObjectID) (models.Task, error) {

	var task models.Task
	err := repo.collection.FindOne(c, bson.M{"_id": id}).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		} else {
			log.Fatal(err)
		}
	}

	return task, nil
}

// Save implements TaskRepository.
func (repo *MongoTaskRepository) Save(c context.Context, task models.Task) (models.Task, error) {

	InsertedTask, err := repo.collection.InsertOne(context.TODO(), task)

	if err != nil {
		return models.Task{}, err
	}

	task.ID = InsertedTask.InsertedID.(primitive.ObjectID)
	return task, nil
}

// Constructor functions

func NewMongoTaskRepository(collection *mongo.Collection) TaskRepository {
	return &MongoTaskRepository{collection}
}
