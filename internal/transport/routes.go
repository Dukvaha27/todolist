package transport

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.Engine,
	taskHandler *TaskHandler,
	categoryHandler *CategoryHandler,
) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("/", taskHandler.List)
		tasks.GET("/:id", taskHandler.GetByID)
		tasks.POST("/", taskHandler.Create)
		tasks.PATCH("/:id", taskHandler.Update)
		tasks.DELETE("/:id", taskHandler.Delete)
	}

	categories := router.Group("/categories")
	{
		categories.GET("/", categoryHandler.List)
		categories.POST("/", categoryHandler.Create)
		categories.DELETE("/:id", categoryHandler.Delete)
	}
}
