package models

import "gorm.io/gorm"

// Category — доменная модель категории
type Category struct {
	gorm.Model
	Name  string  `json:"name" gorm:"not null"`
	Color string  `json:"color" gorm:"default:'#808080'"`
	Tasks []*Task `json:"-" gorm:"foreignKey:CategoryID"`
}

// CategoryCreateRequest — DTO для создания категории
type CategoryCreateRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

// CategoryUpdateRequest — DTO для обновления категории
type CategoryUpdateRequest struct {
	Name  *string `json:"name"`
	Color *string `json:"color"`
}
