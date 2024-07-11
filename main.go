package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Task represents a task with its properties.
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func getTask(c *gin.Context) {
	id := c.Param("id")
	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func removeTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func updateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func addTask(c *gin.Context) {
	var newTask Task

	err := c.ShouldBindJSON(&newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	tasks = append(tasks, newTask)
	c.JSON(http.StatusAccepted, gin.H{"message": "Task created"})
}

func main() {
	fmt.Println("Task Manager API Project")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// Render the HTML template
		c.String(200, "Welcome to Task API")
	})
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTask)
	router.DELETE("/tasks/:id", removeTask)
	router.PUT("/tasks/:id", updateTask)
	router.POST("/tasks", addTask)
	router.Run()
}
