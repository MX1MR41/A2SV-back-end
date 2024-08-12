package data

import (
	"errors"
	"task_manager/models"
)

type ITaskService interface {
	GetTasks() []models.Task
	GetTaskByID(id int) (*models.Task, error)
	CreateTask(task models.Task) models.Task
	UpdateTask(id int, updatedTask models.Task) (*models.Task, error)
	DeleteTask(id int) error
}

type TaskService struct {
	tasks  []models.Task
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: []models.Task{
			{ID: 1, Title: "Task 1", Description: "Description for task 1", DueDate: "2024-08-15", Status: "Pending"},
			{ID: 2, Title: "Task 2", Description: "Description for task 2", DueDate: "2024-08-16", Status: "Completed"},
			{ID: 3, Title: "Task 3", Description: "Description for task 3", DueDate: "2024-08-17", Status: "Pending"},
		},
		nextID: 4,
	}
}

func (s *TaskService) GetTasks() []models.Task {
	return s.tasks
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (s *TaskService) CreateTask(task models.Task) models.Task {
	task.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, task)
	return task
}

func (s *TaskService) UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
	for i, task := range s.tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				task.Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				task.Description = updatedTask.Description
			}
			if updatedTask.DueDate != "" {
				task.DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				task.Status = updatedTask.Status
			}
			s.tasks[i] = task
			return &s.tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func (s *TaskService) DeleteTask(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
