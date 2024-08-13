package Usecases

import (
	"task_manager/Domain"
	"task_manager/Repositories"
)

var taskRepo = Repositories.NewTaskRepository()

type ITaskService interface {
	GetTasks() []Domain.Task
	GetTaskByID(id int) (Domain.Task, error)
	CreateTask(task Domain.Task) (Domain.Task, error)
	UpdateTask(id int, updatedTask Domain.Task) error
	DeleteTask(id int) error
}

type TaskService struct{}

func NewTaskService() ITaskService {
	return &TaskService{}
}

func (t *TaskService) GetTasks() []Domain.Task {

	return taskRepo.GetTasks()
}

func (t *TaskService) GetTaskByID(id int) (Domain.Task, error) {
	task, err := taskRepo.GetTaskByID(id)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (t *TaskService) CreateTask(task Domain.Task) (Domain.Task, error) {
	task.ID = getNextTaskID()

	if err := taskRepo.CreateTask(task); err != nil {
		return task, err
	}
	return task, nil
}

func (t *TaskService) UpdateTask(id int, updatedTask Domain.Task) error {
	_, err := taskRepo.GetTaskByID(id)
	if err != nil {
		return err
	}

	updatedTask.ID = id
	if err := taskRepo.UpdateTask(id, updatedTask); err != nil {
		return err
	}
	return nil
}

func (t *TaskService) DeleteTask(id int) error {
	_, err := taskRepo.GetTaskByID(id)
	if err != nil {
		return err
	}

	if err := taskRepo.DeleteTask(id); err != nil {
		return err
	}
	return nil

}

func getNextTaskID() int {
	return taskRepo.GetNextTaskID()
}
