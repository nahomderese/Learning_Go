package controllers

import (
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-6/task_manager/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskHandlers interface {
	GetAllTasks() gin.HandlerFunc
	GetTaskById() gin.HandlerFunc
	UpdateTask() gin.HandlerFunc
	DeleteTask() gin.HandlerFunc
	CreateTask() gin.HandlerFunc
}

type TaskController struct {
	TaskRepo data.TaskRepository
}

func (ctrl *TaskController) GetAllTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks := ctrl.TaskRepo.FindAll(c)
		c.JSON(200, tasks)
	}
}

func (ctrl *TaskController) GetTaskById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		oid, _ := primitive.ObjectIDFromHex(id)
		task, err := ctrl.TaskRepo.FindByID(c, oid)
		if err != nil {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(200, task)
	}
}

func (ctrl *TaskController) UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		oid, _ := primitive.ObjectIDFromHex(id)
		task, err := ctrl.TaskRepo.FindByID(c, oid)
		if err != nil {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}
		c.BindJSON(&task)
		ctrl.TaskRepo.Save(c, task)
		c.JSON(200, task)
	}
}

func (ctrl *TaskController) DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		oid, _ := primitive.ObjectIDFromHex(id)

		err := ctrl.TaskRepo.Delete(c, oid)
		if err != nil {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(200, gin.H{"message": "Task deleted"})
	}
}

func (ctrl *TaskController) CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		err := c.BindJSON(&task)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctrl.TaskRepo.Save(c, &task)
		c.JSON(200, task)

	}
}
