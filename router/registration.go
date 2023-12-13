package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetRegistrationRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	registrationGroup := r.Group("/")
	{
		registrationGroup.GET("registration", handler.GetRegistration(db))
		registrationGroup.GET("registration/:uuid", handler.GetRegistrationByUUID(db))
		registrationGroup.POST("registration", handler.CreateRegistration(db))
		registrationGroup.PUT("registration/:uuid", handler.UpdateRegistration(db))
		registrationGroup.DELETE("registration/:uuid", handler.DeleteRegistration(db))
		registrationGroup.DELETE("permanent-registration/:uuid", handler.PermanentDeleteRegistration(db))
	}
}
