package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetDoctorRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	doctorGroup := r.Group("/")
	{
		doctorGroup.GET("doctor", handler.GetDoctor(db))
		doctorGroup.GET("doctor/:uuid", handler.GetDoctorByUUID(db))
		doctorGroup.POST("doctor", handler.CreateDoctor(db))
		doctorGroup.PUT("doctor/:uuid", handler.UpdateDoctor(db))
		doctorGroup.DELETE("doctor/:uuid", handler.DeleteDoctor(db))
		doctorGroup.DELETE("permanent-doctor/:uuid", handler.PermanentDeleteDoctor(db))
	}
}
