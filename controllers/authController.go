package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"belajar_go_echo/config"
	"belajar_go_echo/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("my_secret_key")

type JwtCustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// LOGIN
func Login(c echo.Context) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// cek user di DB
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "User not found"})
	}

	// cek password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Password is incorrect"})
	}

	// buat claims
	claims := &JwtCustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// RESTRICTED (protected route)
func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Welcome " + claims.Email,
	})
}

// REGISTER
func Register(c echo.Context) error {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON -> struct
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// Cek email sudah ada atau belum
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Email already registered"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}

	// Simpan user baru
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user"})
	}

	// Buat token JWT langsung setelah register
	// claims := &JwtCustomClaims{
	// 	UserID: user.ID,
	// 	Email:  user.Email,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": echo.Map{
			// "jwt": t,
			"profile": echo.Map{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			},
		},
		"success": true,
	})

}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// generate random token
func generateToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func ForgotPassword(c echo.Context) error {
	fmt.Println(">>> ForgotPassword endpoint called")
	req := new(ForgotPasswordRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// cek apakah user ada
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Email tidak ditemukan"})
	}

	// generate token
	token, err := generateToken(32)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Gagal generate token"})
	}

	// simpan token ke DB
	reset := models.PasswordReset{
		Email:     req.Email,
		Token:     token,
		CreatedAt: time.Now(),
	}
	config.DB.Create(&reset)

	// buat reset link
	resetLink := "http://localhost:3000/reset-password?token=" + token

	// kirim email pakai template
	subject := "Reset Password"
	data := map[string]string{
		"ResetLink": resetLink,
	}

	err = config.SendEmailTemplate(user.Email, subject, "templates/reset_password.html", data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Gagal mengirim email"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Email reset password sudah dikirim",
	})
}

type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

func ResetPassword(c echo.Context) error {
	req := new(ResetPasswordRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// cari token di tabel password_resets
	var reset models.PasswordReset
	if err := config.DB.Where("token = ?", req.Token).First(&reset).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Token tidak valid"})
	}

	// cek apakah token expired (misalnya berlaku 1 jam)
	if time.Since(reset.CreatedAt) > time.Hour {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Token sudah expired"})
	}

	// cari user berdasarkan email
	var user models.User
	if err := config.DB.Where("email = ?", reset.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User tidak ditemukan"})
	}

	// hash password baru
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Gagal hash password"})
	}

	// update password user
	user.Password = string(hashed)
	config.DB.Save(&user)

	// hapus token biar tidak bisa dipakai ulang
	config.DB.Delete(&reset)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Password berhasil direset, silakan login kembali",
	})
}
