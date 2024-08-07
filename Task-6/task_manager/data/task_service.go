package data

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskRepository is an interface for managing tasks.

type TaskRepository interface {
	Save(c context.Context, task *models.Task) (models.Task, error)
	FindByID(c context.Context, id primitive.ObjectID) (*models.Task, error)
	FindAll(c context.Context, user models.UserRes) []models.Task
	Delete(c context.Context, id primitive.ObjectID) error
}

type MongoTaskRepository struct {
	collection *mongo.Collection
}

// Delete implements TaskRepository.
func (repo *MongoTaskRepository) Delete(c context.Context, id primitive.ObjectID) error {

	result, _ := repo.collection.DeleteOne(c, bson.M{"_id": id})

	fmt.Println(result.DeletedCount)

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

// FindAll implements TaskRepository.
func (repo *MongoTaskRepository) FindAll(c context.Context, user models.UserRes) []models.Task {

	query := bson.D{}

	if user.Role != "admin" {
		query = bson.D{{"user_id", user.ID}}
	}

	cursor, err := repo.collection.Find(c, query)

	if err != nil {
		log.Fatal(err)
	}

	var tasks []models.Task
	if err = cursor.All(c, &tasks); err != nil {
		log.Fatal(err)
	}

	for cursor.Next(c) {
		var elem models.Task
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, elem)
	}

	return tasks
}

// FindByID implements TaskRepository.
func (repo *MongoTaskRepository) FindByID(c context.Context, id primitive.ObjectID) (*models.Task, error) {

	var task models.Task
	err := repo.collection.FindOne(c, bson.M{"_id": id}).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		} else {
			log.Fatal(err)
		}
	}

	return &task, nil
}

// Save implements TaskRepository.
func (repo *MongoTaskRepository) Save(c context.Context, task *models.Task) (models.Task, error) {
	task.ID = primitive.NewObjectID()
	_, err := repo.collection.InsertOne(context.TODO(), task)
	if err != nil {
		return models.Task{}, err
	}
	return *task, nil
}

// Constructor functions

func NewMongoTaskRepository(collection *mongo.Collection) TaskRepository {
	return &MongoTaskRepository{collection}
}
