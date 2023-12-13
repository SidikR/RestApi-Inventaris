// handler/registration_handler.go
package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateRegistration(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Registration
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the provided RegisteredId exists in the database
		var patient model.Patient
		if err := db.Where("nik = ?", input.PatientId).First(&patient).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Patient Id not found"})
			return
		}

		// Check if the provided RecipeId exists in the database
		var doctor model.Doctor
		if err := db.Where("uuid = ?", input.DoctorId).First(&doctor).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor not found"})
			return
		}

		// Check if the provided BilId exists in the database
		var poly model.Poly
		if err := db.Where("id = ?", input.PolyId).First(&poly).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bil Report not found"})
			return
		}

		// Save the Registration to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Registration"})
			return
		}

		// Respond with HTTP 201 Created status and the created Registration data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Registration berhasil dibuat",
			Data:    input,
		})
	}
}

func GetRegistration(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.Registration

		// Mengambil data dari tabel "registration" dengan join ke tabel "departemen"
		if err := db.Table("registrations").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data registration"})
			return
		}

		c.JSON(200, result)
	}
}

func GetRegistrationByUUID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var registration struct {
			model.Registration
			Doctor  model.Doctor  `gorm:"foreignkey:DoctorID"`
			Poly    model.Poly    `gorm:"foreignkey:PolyID"`
			Patient model.Patient `gorm:"foreignkey:PatientID"`
		}
		if err := db.Table("registrations").
			Preload("Doctor").
			Preload("Poly").
			Preload("Patient").
			Where("uuid = ?", uuid).
			First(&registration).
			Error; err != nil {
			c.JSON(404, gin.H{"error": "Registration not found"})
			return
		}
		c.JSON(200, registration)
	}
}

func UpdateRegistration(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		uuid := c.Param("uuid")

		// Mencari registration berdasarkan UUID
		var existingRegistration model.Registration
		if err := db.Where("uuid = ?", uuid).First(&existingRegistration).Preload("Doctor").Preload("Poly").Preload("Patient").Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Registration not found"})
			return
		}

		// Binding input JSON ke struct Registration
		var updatedRegistration model.Registration
		if err := c.ShouldBindJSON(&updatedRegistration); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the provided RegisteredId exists in the database
		var patient model.Patient
		if err := db.Where("nik = ?", updatedRegistration.PatientId).First(&patient).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Patient Id not found"})
			return
		}

		// Check if the provided RecipeId exists in the database
		var doctor model.Doctor
		if err := db.Where("uuid = ?", updatedRegistration.DoctorId).First(&doctor).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor not found"})
			return
		}

		// Check if the provided BilId exists in the database
		var poly model.Poly
		if err := db.Where("id = ?", updatedRegistration.PolyId).First(&poly).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bil Report not found"})
			return
		}

		// Memperbarui atribut Registration
		db.Model(&existingRegistration).Updates(&updatedRegistration)

		// Mengambil data Registration yang diperbarui dari database
		var updatedRegistrationFromDB model.Registration
		if err := db.Where("uuid = ?", uuid).First(&updatedRegistrationFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Registration"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Registration successfully updated",
			Data:    updatedRegistrationFromDB,
		})

	}
}

func DeleteRegistration(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var registration model.Registration
		if err := db.Where("uuid = ?", uuid).First(&registration).Error; err != nil {
			c.JSON(404, gin.H{"error": "Registration not found"})
			return
		}

		registrationName := registration.UUID

		db.Delete(&registration)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", registrationName)})
	}
}

func PermanentDeleteRegistration(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var registration model.Registration
		if err := db.Unscoped().Where("UUID = ?", uuid).First(&registration).Error; err != nil {
			c.JSON(404, gin.H{"error": "Registration not found"})
			return
		}

		registrationName := registration.UUID

		db.Unscoped().Delete(&registration)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", registrationName)})
	}
}
