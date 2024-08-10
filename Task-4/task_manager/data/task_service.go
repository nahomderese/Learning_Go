package data

import (
	"errors"

	"github.com/Nahom-Derese/Learning_Go/Task-4/task_manager/models"
)

type TaskRepository interface {
	Save(task *models.Task)
	FindByID(id int64) (*models.Task, error)
	FindAll() []models.Task
	Delete(id int64) error
}

type InMemoryTaskRepository struct {
	idCounter int64
	tasks     map[int64]models.Task
}

type MongoTaskRepository struct {
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		idCounter: 1,
		tasks:     make(map[int64]models.Task),
	}
}

func (repo *InMemoryTaskRepository) Save(task *models.Task) {
	task.ID = repo.idCounter
	repo.idCounter += 1
	repo.tasks[task.ID] = *task
}

func (repo *InMemoryTaskRepository) FindByID(id int64) (*models.Task, error) {
	value, ok := repo.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return &value, nil
}

func (repo *InMemoryTaskRepository) FindAll() []models.Task {
	tasks := make([]models.Task, 0, len(repo.tasks))

	for _, value := range repo.tasks {
		tasks = append(tasks, value)
	}

	return tasks
}

func (repo *InMemoryTaskRepository) Delete(id int64) error {
	_, ok := repo.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(repo.tasks, id)
	return nil

}

type TaskUseCase struct {
	TaskRepo *TaskRepository
}
