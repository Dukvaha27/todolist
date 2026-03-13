package repository

import (
	"todolist-layered/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetByID(id uint) (*models.Category, error)
	List() ([]models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
}

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}
}

func (c *gormCategoryRepository) Create(category *models.Category) error {
	result := c.db.Create(&category)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *gormCategoryRepository) GetByID(id uint) (*models.Category, error) {
	category := models.Category{}

	result := c.db.First(&category, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (c *gormCategoryRepository) List() ([]models.Category, error) {
	categories := []models.Category{}

	result := c.db.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (c *gormCategoryRepository) Update(category *models.Category) error {
	return c.db.Save(&category).Error
}

func (c *gormCategoryRepository) Delete(id uint) error {
	return c.db.Delete(&models.Category{}, id).Error
}
