// package main

// import (
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"github.com/gin-gonic/gin"
// )

// type Task struct {
// 	ID          int    `json:"id"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Status      string `json:"status"`
// }


// var tasks = make(map[int]Task)
// var nextID = 1

// func main() {
// 	r := gin.Default()

// 	// Task Routes
// 	r.POST("/tasks", createTask)
// 	r.GET("/tasks", getTasks)
// 	r.GET("/tasks/:id", getTaskByID)
// 	r.PUT("/tasks/:id", updateTask)
// 	r.DELETE("/tasks/:id", deleteTask)

// 	r.Run(":8080") // Start server on port 8080
// }

// // Create a new task
// func createTask(c *gin.Context) {
// 	var newTask Task
// 	if err := c.ShouldBindJSON(&newTask); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Default status if not provided
// 	if newTask.Status == "" {
// 		newTask.Status = "pending"
// 	}

// 	newTask.ID = nextID
// 	tasks[nextID] = newTask
// 	nextID++

// 	c.JSON(http.StatusCreated, newTask)
// }


// // Get all tasks with pagination, title filtering, and status filtering
// func getTasks(c *gin.Context) {
// 	titleFilter := c.Query("title")
// 	statusFilter := c.Query("status")
// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

// 	if page < 1 {
// 		page = 1
// 	}
// 	if limit < 1 {
// 		limit = 10
// 	}

// 	// Collect matching tasks
// 	var filteredTasks []Task
// 	for _, task := range tasks {
// 		if (titleFilter == "" || containsIgnoreCase(task.Title, titleFilter)) &&
// 			(statusFilter == "" || strings.EqualFold(task.Status, statusFilter)) {
// 			filteredTasks = append(filteredTasks, task)
// 		}
// 	}

// 	// Apply pagination
// 	start := (page - 1) * limit
// 	end := start + limit

// 	if start >= len(filteredTasks) {
// 		c.JSON(http.StatusOK, []Task{}) // Return empty array if out of range
// 		return
// 	}

// 	if end > len(filteredTasks) {
// 		end = len(filteredTasks)
// 	}

// 	c.JSON(http.StatusOK, filteredTasks[start:end])
// }


// // Get task by ID
// func getTaskByID(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 		return
// 	}

// 	task, exists := tasks[id]
// 	if !exists {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, task)
// }

// func updateTask(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 		return
// 	}

// 	var updatedTask Task
// 	if err := c.ShouldBindJSON(&updatedTask); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	existingTask, exists := tasks[id]
// 	if !exists {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
// 		return
// 	}

// 	// Update fields (preserve ID)
// 	existingTask.Title = updatedTask.Title
// 	existingTask.Description = updatedTask.Description
// 	existingTask.Status = updatedTask.Status

// 	tasks[id] = existingTask

// 	c.JSON(http.StatusOK, existingTask)
// }


// // Delete a task
// func deleteTask(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 		return
// 	}

// 	if _, exists := tasks[id]; !exists {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
// 		return
// 	}

// 	delete(tasks, id)
// 	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
// }

// // Helper function: case-insensitive string matching
// func containsIgnoreCase(str, substr string) bool {
// 	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
// }
