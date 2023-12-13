package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetReceipRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	receipGroup := r.Group("/")
	{
		receipGroup.GET("receip", handler.GetReceip(db))
		receipGroup.GET("receip/:uuid", handler.GetReceipByUUID(db))
		receipGroup.POST("receip", handler.CreateReceip(db))
		// receipGroup.PUT("receip/:uuid", handler.UpdateReceip(db))
		receipGroup.DELETE("receip/:uuid", handler.DeleteReceip(db))
		receipGroup.DELETE("permanent-receip/:uuid", handler.PermanentDeleteReceip(db))
	}
}
