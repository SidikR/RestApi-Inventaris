package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetMedicalReportRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	medicalReportGroup := r.Group("/")
	{
		medicalReportGroup.GET("medicalReport", handler.GetMedicalReport(db))
		medicalReportGroup.GET("medicalReport/:uuid", handler.GetMedicalReportByUUID(db))
		medicalReportGroup.DELETE("medicalReport/:uuid", handler.DeleteMedicalReport(db))
		medicalReportGroup.DELETE("permanent-medicalReport/:uuid", handler.PermanentDeleteMedicalReport(db))
	}
}
