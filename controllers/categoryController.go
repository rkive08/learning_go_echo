package controllers

import (
	"belajar_go_echo/config"
	"belajar_go_echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// API
func GetCategories(c echo.Context) error {
	var categories []models.CategoryProduct
	config.DB.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func CreateCategory(c echo.Context) error {
	var category models.CategoryProduct
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	config.DB.Create(&category)
	return c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id") // ambil param ID dari URL

	var category models.CategoryProduct
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	var input models.CategoryProduct
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	category.Name = input.Name
	config.DB.Save(&category)

	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	var category models.CategoryProduct
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	config.DB.Delete(&category)

	return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted"})
}
