package data

import (
	"errors"
	"task_manager/models"
)

var tasks = []models.Task{
	{ID: 1, Title: "Task 1", Description: "Description for task 1", DueDate: "2024-08-15", Status: "Pending"},
	{ID: 2, Title: "Task 2", Description: "Description for task 2", DueDate: "2024-08-16", Status: "Completed"},
	{ID: 3, Title: "Task 3", Description: "Description for task 3", DueDate: "2024-08-17", Status: "Pending"},
}

var nextID = 4

func GetTasks() []models.Task {
	return tasks
}

func GetTaskByID(id int) (*models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(task models.Task) models.Task {
	task.ID = nextID
	nextID++

	tasks = append(tasks, task)
	return task
}

func UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {

			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.DueDate != "" {
				tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}

			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
