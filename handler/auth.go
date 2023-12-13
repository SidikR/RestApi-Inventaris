// handler/auth_handler.go
package handler

import (
	"fmt"
	"net/http"
	"time"

	"main/auth"
	"main/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		input.Password = string(hashedPassword)

		// Set default role (misal: "user")
		input.Role = "user"

		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user model.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials (Username)"})
			return
		}

		// Logging: Tambahkan log untuk melihat role dari pengguna yang berhasil ditemukan
		fmt.Println("User Role:", user.Role)

		// Verifikasi password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials (Password)"})
			return
		}

		// Buat token baru
		token, err := auth.GenerateToken(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// Simpan token di database
		user.Token = token
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// Set cookie dengan nama "token"
		c.SetCookie("token", token, int(time.Hour.Seconds()), "/", "http://localhost:3000", false, true)

		c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful", "username": user.Username})
	}
}

func Logout(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Hapus token di database
		claims, err := auth.VerifyToken(c, db)
		if err == nil {
			var user model.User
			db.Where("username = ?", claims.Username).First(&user)
			user.Token = ""
			db.Save(&user)
		}

		// Hapus cookie
		c.SetCookie("token", "", -1, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	}
}
