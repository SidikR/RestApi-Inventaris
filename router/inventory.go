package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetInventoryRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	inventoryGroup := r.Group("/")
	{
		inventoryGroup.GET("inventory", handler.GetInventory(db))
		inventoryGroup.GET("inventory/:id", handler.GetInventoryByID(db))
		inventoryGroup.POST("inventory", handler.CreateInventory(db))
		inventoryGroup.PUT("inventory/:id", handler.UpdateInventory(db))
		inventoryGroup.DELETE("inventory/:id", handler.DeleteInventory(db))
		inventoryGroup.DELETE("permanent-inventory/:id", handler.PermanentDeleteInventory(db))
	}
}
