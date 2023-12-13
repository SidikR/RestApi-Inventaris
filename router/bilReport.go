package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetBilReportRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	bilReportGroup := r.Group("/")
	{
		bilReportGroup.GET("bilReport", handler.GetBilReport(db))
		bilReportGroup.GET("bilReport/:uuid", handler.GetBilReportByUUID(db))
		// bilReportGroup.POST("bilReport", handler.CreateBilReport(db))
		bilReportGroup.PUT("bilReport/:uuid", handler.UpdateBilReport(db))
		bilReportGroup.DELETE("bilReport/:uuid", handler.DeleteBilReport(db))
		bilReportGroup.DELETE("permanent-bilReport/:uuid", handler.PermanentDeleteBilReport(db))
	}
}
