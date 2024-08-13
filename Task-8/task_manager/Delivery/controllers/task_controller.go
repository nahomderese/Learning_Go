package controllers

import (
	"context"
	"net/http"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
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
	TaskUsecase domain.TaskUsecase
	UserUsecase domain.UserUsecase
}

func NewTaskHandlers(taskUsecase domain.TaskUsecase, userUsecase domain.UserUsecase) TaskHandlers {
	return &TaskController{TaskUsecase: taskUsecase, UserUsecase: userUsecase}
}

func (ctrl *TaskController) GetAllTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(domain.User)
		tasks, err := ctrl.TaskUsecase.FindAll(context.TODO(), user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, tasks)
	}
}

func (ctrl *TaskController) GetTaskById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		oid, _ := primitive.ObjectIDFromHex(id)
		task, err := ctrl.TaskUsecase.FindByID(c, oid)
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
		task, err := ctrl.TaskUsecase.FindByID(c, oid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		c.BindJSON(&task)
		ctrl.TaskUsecase.Save(c, task)
		c.JSON(200, task)
	}
}

func (ctrl *TaskController) DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		oid, _ := primitive.ObjectIDFromHex(id)

		err := ctrl.TaskUsecase.Delete(c, oid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{"message": "Task deleted"})
	}
}

func (ctrl *TaskController) CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task domain.Task
		err := c.BindJSON(&task)

		if err != nil {
			return
		}

		_, e := ctrl.UserUsecase.FindUser(task.UserId)

		if e != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		task, err = ctrl.TaskUsecase.Save(context.TODO(), task)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, task)
	}
}
