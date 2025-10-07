package main

import (
	"belajar_go_echo/config"
	"belajar_go_echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// koneksi database
	config.ConnectDatabase()

	// init router
	e := echo.New()

	// daftarkan routes
	routes.RegisterRoutes(e)

	// jalankan server
	e.Logger.Fatal(e.Start(":8080"))
}
