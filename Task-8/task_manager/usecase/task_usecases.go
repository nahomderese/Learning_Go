package usecase

import (
	"context"
	"errors"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskUseCase is a struct that defines the usecases for the task entity
type TaskUseCase struct {
	TaskRepository domain.TaskRepository
}

// NewTaskUseCase is a function that returns a new TaskUseCase
func NewTaskUseCase(taskRepo domain.TaskRepository) domain.TaskUsecase {
	return &TaskUseCase{TaskRepository: taskRepo}
}

// Delete implements Taskusitory.
func (u *TaskUseCase) Delete(c context.Context, id primitive.ObjectID) error {

	err := u.TaskRepository.Delete(c, id)

	if err != nil {
		return errors.New("task not found")
	}

	return nil
}

// FindAll implements Taskusitory.
func (u *TaskUseCase) FindAll(c context.Context, user domain.User) []domain.Task {

	tasks := u.TaskRepository.FindAll(c, user)
	return tasks
}

// FindByID implements Taskusitory.
func (u *TaskUseCase) FindByID(c context.Context, id primitive.ObjectID) (domain.Task, error) {

	task, err := u.TaskRepository.FindByID(c, id)

	if err != nil {
		return domain.Task{}, errors.New("task not found")
	}

	return task, nil

}

// Save implements Taskusitory.
func (u *TaskUseCase) Save(c context.Context, task domain.Task) (domain.Task, error) {

	task, err := u.TaskRepository.Save(c, task)

	if err != nil {
		return domain.Task{}, errors.New("task not saved")
	}

	return task, nil
}
