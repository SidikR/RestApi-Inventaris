// handler/medicalReport_handler.go
package handler

import (
	"fmt"
	"main/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetMedicalReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []struct {
			model.MedicalReport
			BilReport    model.BilReport    `gorm:"foreignKey:BilId"`
			Registration model.Registration `gorm:"foreignKey:RegisteredId"`
		}

		// Mengambil data dari tabel "medicalReport" dengan join ke tabel "departemen"
		if err := db.Table("medical_reports").Preload("BilReport").Preload("BilReport.Recipe").Preload("Registration").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data medicalReport"})
			return
		}

		c.JSON(200, result)
	}
}

func GetMedicalReportByUUID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var medicalReport []struct {
			model.MedicalReport
			BilReport    model.BilReport    `gorm:"foreignKey:BilReportId"`
			Receip       model.Receip       `gorm:"foreignKey:ReceipId"`
			Registration model.Registration `gorm:"foreignKey:RegisteredId"`
		}
		if err := db.Table("medical_reports").Preload("BilReport").Preload("Receip").Preload("Registration").Where("uuid = ?", uuid).First(&medicalReport).Error; err != nil {
			c.JSON(404, gin.H{"error": "MedicalReport not found"})
			return
		}
		c.JSON(200, medicalReport)
	}
}

func DeleteMedicalReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var medicalReport model.MedicalReport
		if err := db.Where("uuid = ?", uuid).First(&medicalReport).Error; err != nil {
			c.JSON(404, gin.H{"error": "MedicalReport not found"})
			return
		}

		medicalReportName := medicalReport.UUID

		db.Delete(&medicalReport)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", medicalReportName)})
	}
}

func PermanentDeleteMedicalReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var medicalReport model.MedicalReport
		if err := db.Unscoped().Where("UUID = ?", uuid).First(&medicalReport).Error; err != nil {
			c.JSON(404, gin.H{"error": "MedicalReport not found"})
			return
		}

		medicalReportName := medicalReport.UUID

		db.Unscoped().Delete(&medicalReport)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", medicalReportName)})
	}
}
