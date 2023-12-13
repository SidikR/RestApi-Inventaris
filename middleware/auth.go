package middleware

import (
	"main/auth"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AuthMiddleware adalah middleware otentikasi
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verifikasi token
		claims, err := auth.VerifyToken(c, db)
		if err != nil {
			// Hapus cookie jika token tidak valid
			deleteCookie(c, "token")

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Ambil informasi pengguna dari token dan simpan dalam konteks
		user, err := getUserFromClaims(claims, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
			c.Abort()
			return
		}

		// Simpan informasi pengguna dalam konteks
		c.Set("user", user)

		c.Next()
	}
}

// getUserFromClaims mengambil informasi pengguna dari token
func getUserFromClaims(claims *auth.Claims, db *gorm.DB) (*model.User, error) {
	// Lakukan query ke database untuk mendapatkan informasi pengguna berdasarkan ID atau email dari claims
	var user model.User
	if err := db.Where("username = ?", claims.Username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// deleteCookie menghapus cookie dengan nama tertentu
func deleteCookie(c *gin.Context, cookieName string) {
	c.SetCookie(cookieName, "", -1, "/", "", false, true)
}
