package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetInPatientCareRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	inPatientCareGroup := r.Group("/")
	{
		inPatientCareGroup.GET("inPatientCare", handler.GetInPatientCare(db))
		inPatientCareGroup.GET("inPatientCare/:uuid", handler.GetInPatientCareByUUID(db))
		inPatientCareGroup.POST("inPatientCare", handler.CreateInPatientCare(db))
		inPatientCareGroup.PUT("inPatientCare/:uuid", handler.UpdateInPatientCare(db))
		inPatientCareGroup.DELETE("inPatientCare/:uuid", handler.DeleteInPatientCare(db))
		inPatientCareGroup.DELETE("permanent-inPatientCare/:uuid", handler.PermanentDeleteInPatientCare(db))
	}
}
