package data

import (
	"errors"
	"strconv"

	"github.com/Nahom-Derese/Learning_Go/Task-4/task_manager/models"
)

type TaskRepository interface {
	Save(task *models.Task)
	FindByID(id string) (*models.Task, error)
	FindAll() []models.Task
	Delete(id string) error
}

type InMemoryTaskRepository struct {
	idCounter int
	tasks     map[string]models.Task
}

type MongoTaskRepository struct {
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		idCounter: 1,
		tasks:     make(map[string]models.Task),
	}
}

func (repo *InMemoryTaskRepository) Save(task *models.Task) {
	task.ID = strconv.Itoa(repo.idCounter)
	repo.idCounter += 1
	repo.tasks[task.ID] = *task
}

func (repo *InMemoryTaskRepository) FindByID(id string) (*models.Task, error) {
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

func (repo *InMemoryTaskRepository) Delete(id string) error {
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
