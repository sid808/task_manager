package controllers

import (
	"net/http"
	"strconv"
	"task-service/models"
	"task-service/services"
	"task-service/utils"

	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @Summary Create a new task
// @Description Creates a new task with title, description, and status
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task to be created"
// @Success 201 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	createdTask := services.CreateTask(task)
	utils.RespondWithData(c, createdTask)
}

// GetTasks godoc
// @Summary Get all tasks
// @Description Get all tasks with pagination and optional filtering by title and status
// @Tags tasks
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param title query string false "Filter by title"
// @Param status query string false "Filter by status"
// @Success 200 {object} models.GenericResponse
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	title := c.Query("title")
	status := c.Query("status")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	tasks, total := services.GetFilteredTasks(title, status, page, limit)

	response := gin.H{
		"page":  page,
		"limit": limit,
		"total": total,
		"tasks": tasks,
	}

	utils.RespondWithData(c, response)
}

// GetTaskByID godoc
// @Summary Get a task by ID
// @Description Get a specific task using its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Failure 404 {object} models.GenericResponse
// @Router /tasks/{id} [get]
func GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	task, err := services.GetTaskByID(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Task not found")
		return
	}

	utils.RespondWithData(c, task)
}

// UpdateTask godoc
// @Summary Update a task
// @Description Partially update a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body models.Task true "Updated task data"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Failure 404 {object} models.GenericResponse
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := services.UpdateTask(id, updatedTask)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Task not found")
		return
	}

	utils.RespondWithData(c, task)
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Failure 404 {object} models.GenericResponse
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = services.DeleteTask(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Task not found")
		return
	}

	utils.RespondWithData(c, gin.H{"message": "Task deleted"})
}
