package repositories

import "context"

// import "github.com/Nahom-Derese/Learning_Go/Task-7/task_manager/models"

type TaskRepository interface {
	Save(c context.Context, task models.Task) (models.Task, error)
	FindByID(c context.Context, id primitive.ObjectID) (models.Task, error)
	FindAll(c context.Context, user models.User) []models.Task
	Delete(c context.Context, id primitive.ObjectID) error
}

type MongoTaskRepository struct {
	collection *mongo.Collection
}
