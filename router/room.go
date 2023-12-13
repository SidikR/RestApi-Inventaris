package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetRoomRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	roomGroup := r.Group("/")
	{
		roomGroup.GET("room", handler.GetRoom(db))
		roomGroup.GET("room/:id", handler.GetRoomByID(db))
		roomGroup.POST("room", handler.CreateRoom(db))
		roomGroup.PUT("room/:id", handler.UpdateRoom(db))
		roomGroup.DELETE("room/:id", handler.DeleteRoom(db))
		roomGroup.DELETE("permanent-room/:id", handler.PermanentDeleteRoom(db))
	}
}
