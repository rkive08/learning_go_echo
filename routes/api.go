package routes

import (
	"belajar_go_echo/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// category
	e.GET("/api/categories", controllers.GetCategories)
	e.POST("/api/categories", controllers.CreateCategory)
	e.PUT("/api/categories/:id", controllers.UpdateCategory)
	e.DELETE("/api/categories/:id", controllers.DeleteCategory)
}
