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

	// 	// product
	e.GET("/api/products", controllers.GetProducts)
	e.POST("/api/products", controllers.CreateProduct)
	e.PUT("/api/products/:id", controllers.UpdateProduct)
	e.DELETE("/api/products/:id", controllers.DeleteProduct)
}
