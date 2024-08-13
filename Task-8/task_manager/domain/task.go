package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     time.Time          `json:"due_date"`
	Status      string             `json:"status"`
	UserId      string             `bson:"user_id,omitempty" json:"user_id"`
}

//go:generate mockery --name TaskRepository
type TaskRepository interface {
	Save(c context.Context, task Task) (Task, error)
	FindByID(c context.Context, id primitive.ObjectID) (Task, error)
	FindAll(c context.Context, user User) ([]Task, error)
	Delete(c context.Context, id primitive.ObjectID) error
}

//go:generate mockery --name TaskUsecase
type TaskUsecase interface {
	Delete(c context.Context, id primitive.ObjectID) error
	FindAll(c context.Context, user User) ([]Task, error)
	FindByID(c context.Context, id primitive.ObjectID) (Task, error)
	Save(c context.Context, task Task) (Task, error)
}
