package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetMedicineRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	medicineGroup := r.Group("/")
	{
		medicineGroup.GET("medicine", handler.GetMedicine(db))
		medicineGroup.GET("medicine/:uuid", handler.GetMedicineByUUID(db))
		medicineGroup.POST("medicine", handler.CreateMedicine(db))
		medicineGroup.PUT("medicine/:uuid", handler.UpdateMedicine(db))
		medicineGroup.DELETE("medicine/:uuid", handler.DeleteMedicine(db))
		medicineGroup.DELETE("permanent-medicine/:uuid", handler.PermanentDeleteMedicine(db))
	}
}
