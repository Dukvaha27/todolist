package models

import "gorm.io/gorm"

// Task — доменная модель задачи
type Task struct {
	gorm.Model
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	CategoryID  *uint     `json:"category_id"`
	Category    *Category `json:"-"`
}

// TaskCreateRequest — DTO для создания задачи
type TaskCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  *uint  `json:"category_id"`
}

// TaskUpdateRequest — DTO для обновления задачи
// Все поля — указатели, чтобы отличить "не передано" от "передан null"
type TaskUpdateRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
	CategoryID  *uint   `json:"category_id"`
}
