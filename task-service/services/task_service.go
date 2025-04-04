package services

import (
	"errors"
	"task-service/models"
	"strings"
)

var tasks = make(map[int]models.Task)
var nextID = 1

// Create task
func CreateTask(task models.Task) models.Task {
	if task.Status == "" {
		task.Status = "pending"
	}
	task.ID = nextID
	tasks[nextID] = task
	nextID++
	return task
}

// Get all tasks
func GetFilteredTasks(title, status string, page, limit int) ([]models.Task, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	start := (page - 1) * limit
	end := start + limit

	matched := 0
	paginated := []models.Task{}

	for _, task := range tasks {
		if title != "" && !strings.Contains(strings.ToLower(task.Title), strings.ToLower(title)) {
			continue
		}
		if status != "" && strings.ToLower(task.Status) != strings.ToLower(status) {
			continue
		}

		if matched >= start && matched < end {
			paginated = append(paginated, task)
		}
		matched++

		if matched >= end {
			break
		}
	}

	return paginated, matched
}

// Get task by ID
func GetTaskByID(id int) (models.Task, error) {
	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

// Update task
func UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
	if _, exists := tasks[id]; !exists {
		return models.Task{}, errors.New("task not found")
	}

	updatedTask.ID = id
	tasks[id] = updatedTask
	return updatedTask, nil
}

// Delete task
func DeleteTask(id int) error {
	if _, exists := tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}
