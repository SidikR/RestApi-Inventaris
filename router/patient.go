package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetPatientRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	patientGroup := r.Group("/")
	{
		patientGroup.GET("patient", handler.GetPatient(db))
		patientGroup.GET("patient/:nik", handler.GetPatientByNIK(db))
		patientGroup.POST("patient", handler.CreatePatient(db))
		patientGroup.PUT("patient/:nik", handler.UpdatePatient(db))
		patientGroup.DELETE("patient/:nik", handler.DeletePatient(db))
		patientGroup.DELETE("permanent-patient/:nik", handler.PermanentDeletePatient(db))
	}
}
