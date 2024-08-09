package controllers

import (
	"context"
	"net/http"

	"github.com/Nahom-Derese/Learning_Go/Task-7/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-7/task_manager/models"
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
	UserRepo data.UserRepository
}

func (ctrl *TaskController) GetAllTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)
		tasks := ctrl.TaskRepo.FindAll(context.TODO(), user)
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
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
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
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "Task deleted"})
	}
}

func (ctrl *TaskController) CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		err := c.BindJSON(&task)

		if err != nil {
			return
		}

		_, e := ctrl.UserRepo.FindUser(task.UserId)

		if e != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		task, err = ctrl.TaskRepo.Save(context.TODO(), task)
		c.JSON(200, task)

	}
}
