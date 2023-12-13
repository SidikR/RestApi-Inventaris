package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetPolyRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	polyGroup := r.Group("/")
	{
		polyGroup.GET("poly", handler.GetPoly(db))
		polyGroup.GET("poly/:id", handler.GetPolyByID(db))
		polyGroup.POST("poly", handler.CreatePoly(db))
		polyGroup.PUT("poly/:id", handler.UpdatePoly(db))
		polyGroup.DELETE("poly/:id", handler.DeletePoly(db))
		polyGroup.DELETE("permanent-poly/:id", handler.PermanentDeletePoly(db))
	}
}
