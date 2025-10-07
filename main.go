package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Membuat instance Echo
	e := echo.New()

	// Route GET sederhana
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Halo dari Echo!")
	})

	// Route lain (contoh dengan JSON)
	e.GET("/user", func(c echo.Context) error {
		user := map[string]string{
			"name":  "Ratna",
			"email": "ratna@example.com",
		}
		return c.JSON(http.StatusOK, user)
	})

	// Menjalankan server di port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
