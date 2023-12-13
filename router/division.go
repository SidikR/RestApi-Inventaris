package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetDivisionRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	divisionGroup := r.Group("/")
	{
		divisionGroup.GET("division", handler.GetDivision(db))
		divisionGroup.GET("division/:id", handler.GetDivisionByID(db))
		divisionGroup.POST("division", handler.CreateDivision(db))
		divisionGroup.PUT("division/:id", handler.UpdateDivision(db))
		divisionGroup.DELETE("division/:id", handler.DeleteDivision(db))
		divisionGroup.DELETE("permanent-division/:id", handler.PermanentDeleteDivision(db))
	}
}
