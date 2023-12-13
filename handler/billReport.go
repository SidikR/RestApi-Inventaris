// handler/bilReport_handler.go
package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetBilReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []struct {
			model.BilReport
			Receip       model.Receip       `gorm:"foreignKey:RecipeId"`
			Registration model.Registration `gorm:"foreignKey:RegisteredId"`
		}
		db.LogMode(true)

		// Mengambil data dari tabel "bilReport" dengan join ke tabel "departemen"
		if err := db.Table("bil_reports").Preload("Receip").Preload("Registration").Preload("Registration.Patient").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data bilReport"})
			return
		}

		c.JSON(200, result)
	}
}

func GetBilReportByUUID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var bilReport model.BilReport

		// Preload data dengan relasi yang bersarang
		if err := db.Preload("Receip").
			Preload("Registration", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Patient")
			}).
			Where("uuid = ?", uuid).
			First(&bilReport).Error; err != nil {
			c.JSON(404, gin.H{"error": "Bil Report not found"})
			return
		}

		c.JSON(200, bilReport)
	}
}

func UpdateBilReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		uuid := c.Param("uuid")

		// Mencari bilReport berdasarkan UUID
		var existingBilReport model.BilReport
		if err := db.Where("uuid = ?", uuid).First(&existingBilReport).Preload("InPatientCare").Preload("Receip").Preload("Registration").Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "BilReport not found"})
			return
		}

		// Binding input JSON ke struct BilReport
		var updatedBilReport model.BilReport
		if err := c.ShouldBindJSON(&updatedBilReport); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the provided RegisteredId exists in the database
		var registration model.Registration
		if err := db.Where("id = ?", updatedBilReport.RegisteredId).First(&registration).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Registration Id not found"})
			return
		}

		// Check if the provided RecipeId exists in the database
		var receip model.Receip
		if err := db.Where("id = ?", updatedBilReport.RecipeId).First(&receip).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Receip not found"})
			return
		}

		// Memperbarui atribut BilReport
		db.Model(&existingBilReport).Updates(&updatedBilReport)

		// Mengambil data BilReport yang diperbarui dari database
		var updatedBilReportFromDB model.BilReport
		if err := db.Where("uuid = ?", uuid).First(&updatedBilReportFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated BilReport"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "BilReport successfully updated",
			Data:    updatedBilReportFromDB,
		})

	}
}

func DeleteBilReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var bilReport model.BilReport
		if err := db.Where("uuid = ?", uuid).First(&bilReport).Error; err != nil {
			c.JSON(404, gin.H{"error": "BilReport not found"})
			return
		}

		bilReportName := bilReport.UUID

		db.Delete(&bilReport)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", bilReportName)})
	}
}

func PermanentDeleteBilReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var bilReport model.BilReport
		if err := db.Unscoped().Where("UUID = ?", uuid).First(&bilReport).Error; err != nil {
			c.JSON(404, gin.H{"error": "Bil Report not found"})
			return
		}

		db.Unscoped().Delete(&bilReport)
		c.JSON(200, gin.H{"message": fmt.Sprintf("Permanently deleted")})
	}
}
