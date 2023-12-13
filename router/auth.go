package router

import (
	"main/handler"
	"main/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetAuthRoutes(r *gin.RouterGroup, db *gorm.DB) {
	authMiddleware := middleware.AuthMiddleware(db)

	// Register endpoint
	r.POST("/register", handler.Register(db))

	// Login endpoint
	r.POST("/login", handler.Login(db))

	// Logout endpoint with authentication middleware
	r.POST("/logout", authMiddleware, handler.Logout(db))

	authGroup := r.Group("/auth", authMiddleware)
	{
		// Example endpoint that requires authentication
		authGroup.GET("/profile", func(c *gin.Context) {
			// Get user information from context
			user, exists := c.Get("user")
			if !exists {
				c.JSON(500, gin.H{"error": "Internal Server Error"})
				return
			}

			// Use user information
			c.JSON(200, gin.H{"message": "Profile", "user": user})
		})
	}
}
