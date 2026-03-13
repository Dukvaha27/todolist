package repository

import (
	"todolist-layered/internal/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
	List() ([]models.Task, error)
}

type gormTaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &gormTaskRepository{db: db}
}

func (r *gormTaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *gormTaskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *gormTaskRepository) Update(task *models.Task) error {
	result := r.db.Save(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *gormTaskRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Task{}, id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *gormTaskRepository) List() ([]models.Task, error) {
	tasks := []models.Task{}
	result := r.db.Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
