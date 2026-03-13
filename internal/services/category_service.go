package services

import (
	"todolist-layered/internal/models"
	"todolist-layered/internal/repository"
)

type CategoryService interface {
	GetByID(id uint) (*models.Category, error)
	List() ([]models.Category, error)
	Update(req *models.CategoryUpdateRequest) (*models.Category, error)
	Create(req *models.CategoryCreateRequest) (*models.Category, error)
	Delete(id uint) error
}

type categoryService struct {
	taskRepo     repository.TaskRepository
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository, taskRepo repository.TaskRepository) CategoryService {
	return &categoryService{
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
	}
}

func (c *categoryService) GetByID(id uint) (*models.Category, error) {
	return c.categoryRepo.GetByID(id)
}

func (c *categoryService) List() ([]models.Category, error) {
	return c.categoryRepo.List()
}

func (c *categoryService) Update(req *models.CategoryUpdateRequest) (*models.Category, error) {
	category := models.Category{}

	if req.Color != nil {
		category.Color = *req.Color
	}

	if req.Name != nil {
		category.Name = *req.Name
	}

	return &category, c.categoryRepo.Update(req)
}

func (c *categoryService) Create(req *models.CategoryCreateRequest) (*models.Category, error) {
	category := models.Category{
		Color: req.Color,
		Name:  req.Name,
	}

	return &category, c.categoryRepo.Create(&category)
}

func (c *categoryService) Delete(id uint) error {
	return c.categoryRepo.Delete(id)
}
