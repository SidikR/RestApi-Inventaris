// handler/inPatientCare_handler.go
package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateInPatientCare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.InPatientCare
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the provided RegisteredId exists in the database
		var registration model.Registration
		if err := db.Where("id = ?", input.RegisteredId).First(&registration).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Registration Id not found"})
			return
		}

		// Check if the provided RecipeId exists in the database
		var room model.Room
		if err := db.Where("id = ?", input.RoomId).First(&room).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
			return
		}

		// Save the InPatientCare to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create InPatientCare"})
			return
		}

		// Respond with HTTP 201 Created status and the created InPatientCare data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "InPatientCare berhasil dibuat",
			Data:    input,
		})
	}
}

func GetInPatientCare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.InPatientCare

		// Mengambil data dari tabel "inPatientCare" dengan join ke tabel "departemen"
		if err := db.Preload("InPatientCare").Preload("Registration").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data inPatientCare"})
			return
		}

		c.JSON(200, result)
	}
}

func GetInPatientCareByUUID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var inPatientCare model.InPatientCare
		if err := db.Where("uuid = ?", uuid).First(&inPatientCare).Preload("InPatientCare").Preload("Registration").Error; err != nil {
			c.JSON(404, gin.H{"error": "InPatientCare not found"})
			return
		}
		c.JSON(200, inPatientCare)
	}
}

func UpdateInPatientCare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		uuid := c.Param("uuid")

		// Mencari inPatientCare berdasarkan UUID
		var existingInPatientCare model.InPatientCare
		if err := db.Where("uuid = ?", uuid).First(&existingInPatientCare).Preload("Room").Preload("Registration").Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "InPatientCare not found"})
			return
		}

		// Binding input JSON ke struct InPatientCare
		var updatedInPatientCare model.InPatientCare
		if err := c.ShouldBindJSON(&updatedInPatientCare); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the provided RegisteredId exists in the database
		var registration model.Registration
		if err := db.Where("id = ?", updatedInPatientCare.RegisteredId).First(&registration).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Registration Id not found"})
			return
		}

		// Check if the provided RecipeId exists in the database
		// Check if the provided RecipeId exists in the database
		var room model.Room
		if err := db.Where("id = ?", updatedInPatientCare.RoomId).First(&room).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
			return
		}

		// Memperbarui atribut InPatientCare
		db.Model(&existingInPatientCare).Updates(&updatedInPatientCare)

		// Mengambil data InPatientCare yang diperbarui dari database
		var updatedInPatientCareFromDB model.InPatientCare
		if err := db.Where("uuid = ?", uuid).First(&updatedInPatientCareFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated InPatientCare"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "InPatientCare successfully updated",
			Data:    updatedInPatientCareFromDB,
		})

	}
}

func DeleteInPatientCare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var inPatientCare model.InPatientCare
		if err := db.Where("uuid = ?", uuid).First(&inPatientCare).Error; err != nil {
			c.JSON(404, gin.H{"error": "InPatientCare not found"})
			return
		}

		inPatientCareName := inPatientCare.UUID

		db.Delete(&inPatientCare)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", inPatientCareName)})
	}
}

func PermanentDeleteInPatientCare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var inPatientCare model.InPatientCare
		if err := db.Unscoped().Where("UUID = ?", uuid).First(&inPatientCare).Error; err != nil {
			c.JSON(404, gin.H{"error": "Patient not found"})
			return
		}

		db.Unscoped().Delete(&inPatientCare)
		c.JSON(200, gin.H{"message": fmt.Sprintf("Permanently deleted")})
	}
}
