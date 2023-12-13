package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Patient
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the Patient to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Patient"})
			return
		}

		// Respond with HTTP 201 Created status and the created Patient data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Patient created succesfully",
			Data:    input,
		})
	}
}

func GetPatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.Patient

		// Mengambil data dari tabel "patient" dengan join ke tabel "patient"
		if err := db.Table("patients").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Not Found Patient"})
			return
		}

		c.JSON(200, result)
	}
}

func GetPatientByNIK(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		nik := c.Param("nik")
		var patient model.Patient
		if err := db.Where("NIK = ?", nik).First(&patient).Error; err != nil {
			c.JSON(404, gin.H{"error": "Patient not found"})
			return
		}
		c.JSON(200, patient)
	}
}

func UpdatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		nik := c.Param("nik")

		// Mencari patient berdasarkan UUID
		var existingPatient model.Patient
		if err := db.Where("NIK = ?", nik).First(&existingPatient).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		// Binding input JSON ke struct Patient
		var updatedPatient model.Patient
		if err := c.ShouldBindJSON(&updatedPatient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Memperbarui atribut Patient
		db.Model(&existingPatient).Updates(&updatedPatient)

		// Mengambil data Patient yang diperbarui dari database
		var updatedPatientFromDB model.Patient
		if err := db.Where("nik = ?", nik).First(&updatedPatientFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Patient"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Patient successfully updated",
			Data:    updatedPatientFromDB,
		})

	}
}

func DeletePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		nik := c.Param("nik")
		var patient model.Patient
		if err := db.Where("NIK = ?", nik).First(&patient).Error; err != nil {
			c.JSON(404, gin.H{"error": "Patient not found"})
			return
		}

		patientName := patient.PatientName

		db.Delete(&patient)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", patientName)})
	}
}

func PermanentDeletePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		nik := c.Param("nik")
		var patient model.Patient
		if err := db.Unscoped().Where("NIK = ?", nik).First(&patient).Error; err != nil {
			c.JSON(404, gin.H{"error": "Patient not found"})
			return
		}

		patientName := patient.PatientName

		db.Unscoped().Delete(&patient)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", patientName)})
	}
}
