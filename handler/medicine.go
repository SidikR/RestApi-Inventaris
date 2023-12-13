package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateMedicine(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Medicine
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the Medicine to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Medicine"})
			return
		}

		// Respond with HTTP 201 Created status and the created Medicine data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Medicine created succesfully",
			Data:    input,
		})
	}
}

func GetMedicine(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.Medicine

		// Mengambil data dari tabel "medicine" dengan join ke tabel "department"
		if err := db.Table("medicines").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data obat"})
			return
		}

		c.JSON(200, result)
	}
}

func GetMedicineByUUID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var medicine model.Medicine
		if err := db.Where("uuid = ?", uuid).First(&medicine).Error; err != nil {
			c.JSON(404, gin.H{"error": "Medicine not found"})
			return
		}
		c.JSON(200, medicine)
	}
}

func UpdateMedicine(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		uuid := c.Param("uuid")

		// Mencari medicine berdasarkan UUID
		var existingMedicine model.Medicine
		if err := db.Where("uuid = ?", uuid).First(&existingMedicine).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
			return
		}

		// Binding input JSON ke struct Medicine
		var updatedMedicine model.Medicine
		if err := c.ShouldBindJSON(&updatedMedicine); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Memperbarui atribut Medicine
		db.Model(&existingMedicine).Updates(&updatedMedicine)

		// Mengambil data Medicine yang diperbarui dari database
		var updatedMedicineFromDB model.Medicine
		if err := db.Where("uuid = ?", uuid).First(&updatedMedicineFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Medicine"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Medicine successfully updated",
			Data:    updatedMedicineFromDB,
		})

	}
}

func DeleteMedicine(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var medicine model.Medicine
		if err := db.Where("uuid = ?", uuid).First(&medicine).Error; err != nil {
			c.JSON(404, gin.H{"error": "Medicine not found"})
			return
		}

		medicineName := medicine.MedicineName

		db.Delete(&medicine)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", medicineName)})
	}
}

func PermanentDeleteMedicine(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var medicalReport model.Medicine
		if err := db.Unscoped().Where("UUID = ?", uuid).First(&medicalReport).Error; err != nil {
			c.JSON(404, gin.H{"error": "Medicine not found"})
			return
		}

		medicalReportName := medicalReport.MedicineName

		db.Unscoped().Delete(&medicalReport)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", medicalReportName)})
	}
}
