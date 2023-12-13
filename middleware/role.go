// middleware/role_middleware.go
package middleware

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware adalah middleware otorisasi berdasarkan role
func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Role Middleware Start")

		user, exists := c.Get("user")
		if !exists {
			fmt.Println("User not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		// Konversi informasi user ke dalam bentuk model.User
		userModel, ok := user.(*model.User)
		if !ok {
			fmt.Println("User conversion failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		// Cek apakah user memiliki role yang diizinkan
		if role != "" && userModel.Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden", "message": "User role does not match the required role"})
			c.Abort()
			return
		}

		c.Next()
	}
}
