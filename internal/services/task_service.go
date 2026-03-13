package services

import (
	"fmt"
	"todolist-layered/internal/models"
	"todolist-layered/internal/repository"
)

type TaskService interface {
	CreateTask(req models.TaskCreateRequest) (*models.Task, error)
	GetTask(id uint) (*models.Task, error)
	UpdateTask(id uint, req models.TaskUpdateRequest) (*models.Task, error)
	DeleteTask(id uint) error
	ListTasks() ([]models.Task, error)
}

type taskService struct {
	taskRepo     repository.TaskRepository
	categoryRepo repository.CategoryRepository
}

func NewTaskService(
	taskRepo repository.TaskRepository,
	categoryRepo repository.CategoryRepository,
) TaskService {
	return &taskService{
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
	}
}

func (t *taskService) CreateTask(req models.TaskCreateRequest) (*models.Task, error) {
	taskCreate := models.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	if req.CategoryID != nil {
		category, err := t.categoryRepo.GetByID(*req.CategoryID)
		if err == nil {
			taskCreate.CategoryID = &category.ID
		} else {
			return nil, err
		}
	}

	err := t.taskRepo.Create(&taskCreate)

	if err != nil {
		return nil, err
	}

	task, err := t.taskRepo.GetByID(taskCreate.ID)

	return task, err

}

func (t *taskService) GetTask(id uint) (*models.Task, error) {
	return t.taskRepo.GetByID(id)
}

func (t *taskService) UpdateTask(id uint, req models.TaskUpdateRequest) (*models.Task, error) {
	task, err := t.taskRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("Task is not found: %s", err)
	}

	if req.Title != nil {
		task.Title = *req.Title
	}

	if req.CategoryID != nil {
		task.CategoryID = req.CategoryID
	}

	if req.Completed != nil {
		task.Completed = *req.Completed
	}

	if req.Description != nil {
		task.Description = *req.Description
	}

	err = t.taskRepo.Update(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskService) DeleteTask(id uint) error {
	return t.taskRepo.Delete(id)
}

func (t *taskService) ListTasks() ([]models.Task, error) {
	return t.taskRepo.List()
}
