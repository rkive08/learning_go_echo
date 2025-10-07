package main

import (
	"belajar_go_echo/config"
	"belajar_go_echo/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// koneksi database
	config.ConnectDatabase()

	// init router
	e := echo.New()

	// ✅ Tambahkan middleware CORS di sini
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"}, // asal Nuxt kamu
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	// daftarkan routes
	routes.RegisterRoutes(e)

	// jalankan server
	e.Logger.Fatal(e.Start(":8080"))
}
