package routes

import (
	"belajar_go_echo/controllers"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// auth
	e.POST("/api/login", controllers.Login)
	e.POST("/api/register", controllers.Register)
	e.POST("/api/forgot-password", controllers.ForgotPassword)
	e.POST("/api/reset-password", controllers.ResetPassword)

	// route yang butuh token
	r := e.Group("/api")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("my_secret_key"),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(controllers.JwtCustomClaims)
		},
	}))

	r.GET("/restricted", controllers.Restricted)

	// 	// category
	r.GET("/category-products", controllers.GetCategories)
	e.POST("/api/category-products", controllers.CreateCategory)
	e.PUT("/api/category-products/:id", controllers.UpdateCategory)
	e.DELETE("/api/category-products/:id", controllers.DeleteCategory)

	// product
	e.GET("/api/products", controllers.GetProducts)
	e.POST("/api/products", controllers.CreateProduct)
	e.PUT("/api/products/:id", controllers.UpdateProduct)
	e.DELETE("/api/products/:id", controllers.DeleteProduct)

	// just test email
	// e.GET("/api/test-email", func(c echo.Context) error {
	// 	err := config.SendEmail("test@example.com", "Hello from Go", "Ini body email percobaan.")
	// 	if err != nil {
	// 		return c.JSON(500, map[string]string{"message": "gagal kirim email"})
	// 	}
	// 	return c.JSON(200, map[string]string{"message": "email terkirim"})
	// })
	// end test email

	// =================================
	// for list all routes on terminal
	// for _, r := range e.Routes() {
	// 	println(r.Method, r.Path)
	// }
	// =================================

}

// func RegisterRoutes(e *echo.Echo) {
// 	// category
// 	e.GET("/api/category-products", controllers.GetCategories)
// 	e.POST("/api/category-products", controllers.CreateCategory)
// 	e.PUT("/api/category-products/:id", controllers.UpdateCategory)
// 	e.DELETE("/api/category-products/:id", controllers.DeleteCategory)

// 	// product
// 	e.GET("/api/products", controllers.GetProducts)
// 	e.POST("/api/products", controllers.CreateProduct)
// 	e.PUT("/api/products/:id", controllers.UpdateProduct)
// 	e.DELETE("/api/products/:id", controllers.DeleteProduct)

// }
