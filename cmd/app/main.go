package main

import (
	"todolist-layered/internal/config"
	"todolist-layered/internal/models"
	"todolist-layered/internal/repository"
	"todolist-layered/internal/services"
	"todolist-layered/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDatabase()

	db.AutoMigrate(&models.Task{}, &models.Category{})

	taskRepo := repository.NewTaskRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	taskService := services.NewTaskService(taskRepo, categoryRepo)
	categoryService := services.NewCategoryService(categoryRepo)

	taskHandler := transport.NewTaskHandler(taskService)
	categoryHandler := transport.NewCategoryHandler(categoryService)

	r := gin.Default()

	transport.RegisterRoutes(r, taskHandler, categoryHandler)

	r.Run(":3000")
}
