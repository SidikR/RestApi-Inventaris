package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateDoctor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Doctor
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the Doctor to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Doctor"})
			return
		}

		// Respond with HTTP 201 Created status and the created Doctor data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Doctor created succesfully",
			Data:    input,
		})
	}
}

func GetDoctor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.Doctor

		// Mengambil data dari tabel "doctor" dengan join ke tabel "doctor"
		if err := db.Table("doctors").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Not Found Doctor"})
			return
		}

		c.JSON(200, result)
	}
}

func GetDoctorByUUID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var doctor model.Doctor
		if err := db.Where("uuid = ?", uuid).First(&doctor).Error; err != nil {
			c.JSON(404, gin.H{"error": "Doctor not found"})
			return
		}
		c.JSON(200, doctor)
	}
}

func UpdateDoctor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		uuid := c.Param("uuid")

		// Mencari doctor berdasarkan UUID
		var existingDoctor model.Doctor
		if err := db.Where("uuid = ?", uuid).First(&existingDoctor).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
			return
		}

		// Binding input JSON ke struct Doctor
		var updatedDoctor model.Doctor
		if err := c.ShouldBindJSON(&updatedDoctor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Memperbarui atribut Doctor
		db.Model(&existingDoctor).Updates(&updatedDoctor)

		// Mengambil data Doctor yang diperbarui dari database
		var updatedDoctorFromDB model.Doctor
		if err := db.Where("uuid = ?", uuid).First(&updatedDoctorFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Doctor"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Doctor successfully updated",
			Data:    updatedDoctorFromDB,
		})

	}
}

func DeleteDoctor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var doctor model.Doctor
		if err := db.Where("uuid = ?", uuid).First(&doctor).Error; err != nil {
			c.JSON(404, gin.H{"error": "Doctor not found"})
			return
		}

		doctorName := doctor.DoctorName

		db.Delete(&doctor)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", doctorName)})
	}
}

func PermanentDeleteDoctor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var doctor model.Doctor
		if err := db.Unscoped().Where("uuid = ?", uuid).First(&doctor).Error; err != nil {
			c.JSON(404, gin.H{"error": "Doctor not found"})
			return
		}

		doctorName := doctor.DoctorName

		db.Unscoped().Delete(&doctor)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", doctorName)})
	}
}
