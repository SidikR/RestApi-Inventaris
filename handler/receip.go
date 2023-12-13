// handler/receip_handler.go
package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateReceip(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var receipt model.Receip

		if err := c.ShouldBindJSON(&receipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
			return
		}

		for _, medicine := range receipt.ListMedicines {
			var medicineModel model.Medicine
			if err := db.Where("uuid = ?", medicine.MedicineID).First(&medicineModel).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch medicine", "details": err.Error()})
				return
			}
			medicineModel.Stock -= medicine.Qty
			receipt.TotalPrice += medicine.Qty * medicineModel.Price
		}

		tx := db.Begin()

		// Create Receipt
		if err := tx.Create(&receipt).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Receipt", "details": err.Error()})
			return
		}

		// Create BilReport
		bilReport := model.BilReport{
			RegisteredId: receipt.RegisteredId,
			RecipeId:     receipt.UUID,
		}

		var rates model.Rates
		if err := db.First(&rates).Error; err != nil {
			bilReport.Services = 100000
		} else {
			bilReport.Services = rates.Service
		}

		// Calculate the Total field based on your requirements
		bilReport.Total = bilReport.Services + receipt.TotalPrice

		if err := tx.Create(&bilReport).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create BilReport", "details": err.Error()})
			return
		}

		medicineReport := model.MedicalReport{
			RegisteredId: receipt.RegisteredId,
			BilId:        bilReport.UUID,
		}

		if err := tx.Create(&medicineReport).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Medicine Report", "details": err.Error()})
			return
		}

		tx.Commit()

		c.JSON(http.StatusCreated, receipt)
	}
}

func GetReceip(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []struct {
			model.Receip
			Medicines    []model.ReceiptMedicine `gorm:"foreignKey:ReceipUUID"`
			Registration model.Registration      `gorm:"foreignKey:RegisteredId;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
		}

		if err := db.Table("receips").Preload("Medicines").Preload("Registration").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data receip"})
			return
		}

		c.JSON(200, result)
	}
}

func GetReceipByUUID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var receip []struct {
			model.Receip
			Medicines    []model.ReceiptMedicine `gorm:"foreignKey:ReceipUUID"`
			Registration model.Registration      `gorm:"foreignKey:RegisteredId;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
		}

		uuid := c.Param("uuid")

		if err := db.Table("receips").Preload("Medicines").Preload("Registration").Where("uuid = ?", uuid).First(&receip).Error; err != nil {
			c.JSON(404, gin.H{"error": "Receip not found"})
			return
		}
		c.JSON(200, receip)
	}
}

func DeleteReceip(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var receip model.Receip
		if err := db.Where("uuid = ?", uuid).First(&receip).Error; err != nil {
			c.JSON(404, gin.H{"error": "Receip not found"})
			return
		}

		// Delete associated receipt medicines
		if err := db.Where("receip_uuid = ?", uuid).Delete(model.ReceiptMedicine{}).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete associated receipt medicines"})
			return
		}

		receipName := receip.UUID
		db.Delete(&receip)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", receipName)})
	}
}

func PermanentDeleteReceip(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var receip model.Receip
		if err := db.Unscoped().Where("UUID = ?", uuid).First(&receip).Error; err != nil {
			c.JSON(404, gin.H{"error": "Receip not found"})
			return
		}

		// Permanently delete associated receipt medicines
		if err := db.Where("receip_uuid = ?", uuid).Delete(model.ReceiptMedicine{}).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to permanently delete associated receipt medicines"})
			return
		}

		receipName := receip.UUID
		db.Unscoped().Delete(&receip)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", receipName)})
	}
}
