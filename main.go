// ini tanpa env
// package main

// import (
// 	"belajar_go_echo/config"
// 	"belajar_go_echo/routes"

// 	// seeders "belajar_go_echo/seeder"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	// koneksi database
// 	config.ConnectDatabase()

// 	// auto migrate semua model
// 	// config.DB.AutoMigrate(&models.CategoryProduct{})

// 	// seeders.SeedUser()

// 	// init router
// 	e := echo.New()

// 	// daftarkan routes
// 	routes.RegisterRoutes(e)

// 	// jalankan server
// 	e.Logger.Fatal(e.Start(":8080"))
// }
// end tanpa env

package main

import (
	"belajar_go_echo/config"
	"belajar_go_echo/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system env")
	}

	// koneksi database
	config.ConnectDatabase()

	// load mail config
	config.LoadMailConfig()

	e := echo.New()

	// ✅ Tambahkan middleware CORS di sini
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"}, // asal Nuxt kamu
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(e)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	e.Logger.Fatal(e.Start(":" + port))
}
