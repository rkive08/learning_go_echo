package controllers

import (
	"belajar_go_echo/config"
	"belajar_go_echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// API

func GetProducts(c echo.Context) error {
	var products []models.Product
	if err := config.DB.Preload("Category").Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	var product models.Product

	// Bind JSON ke struct
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Simpan ke DB
	if err := config.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Ambil lagi dengan preload category
	var createdProduct models.Product
	if err := config.DB.Preload("Category").First(&createdProduct, product.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdProduct)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id") // ambil param ID dari URL

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	var input models.Product
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	product.Name = input.Name
	product.Price = input.Price
	product.CategoryID = input.CategoryID

	if err := config.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// reload dengan relasi Category
	config.DB.Preload("Category").First(&product, product.ID)

	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}
