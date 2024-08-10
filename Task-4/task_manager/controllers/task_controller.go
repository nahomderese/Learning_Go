package controllers

import (
	"net/http"
	"strconv"

	"github.com/Nahom-Derese/Learning_Go/Task-4/task_manager/data"
	"github.com/Nahom-Derese/Learning_Go/Task-4/task_manager/models"
	"github.com/gin-gonic/gin"
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
		tasks := ctrl.TaskRepo.FindAll()
		c.JSON(http.StatusOK, tasks)
	}
}

func (ctrl *TaskController) GetTaskById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// convert string to int64
		id_n, err := strconv.Atoi(id)
		id_num := int64(id_n)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		task, err := ctrl.TaskRepo.FindByID(id_num)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusOK, task)
	}
}

func (ctrl *TaskController) UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// convert string to int64
		id_n, err := strconv.Atoi(id)
		id_num := int64(id_n)
		task, err := ctrl.TaskRepo.FindByID(id_num)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.BindJSON(&task)
		ctrl.TaskRepo.Save(task)
		c.JSON(http.StatusOK, task)
	}
}

func (ctrl *TaskController) DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// convert string to int64
		id_n, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}
		id_num := int64(id_n)

		e := ctrl.TaskRepo.Delete(id_num)

		if e != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{"message": "Task deleted"})
	}
}

func (ctrl *TaskController) CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		err := c.BindJSON(&task)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctrl.TaskRepo.Save(&task)
		c.JSON(http.StatusOK, task)

	}
}
