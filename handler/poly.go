// handler/poly_handler.go
package handler

import (
	"fmt"
	"main/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreatePoly(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Poly
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the provided IDRoom exists in the database
		var room model.Room
		if err := db.Where("id = ?", input.RoomId).First(&room).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
			return
		}

		// Save the Poly to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Poly"})
			return
		}

		// Respond with HTTP 201 Created status and the created Poly data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Poly berhasil dibuat",
			Data:    input,
		})
	}
}

func GetPoly(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.Poly

		// Mengambil data dari tabel "polys" dengan preload ke tabel "rooms" dan "doctors"
		if err := db.Table("polies").Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Poly not found"})
			return
		}

		c.JSON(200, result)
	}
}

func GetPolyByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var poly model.Poly
		if err := db.Table("polies").Where("id = ?", id).First(&poly).Error; err != nil {
			c.JSON(404, gin.H{"error": "Poly not found"})
			return
		}
		c.JSON(200, poly)
	}
}
func UpdatePoly(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		id := c.Param("id")

		// Mencari poly berdasarkan UUID
		var existingPoly model.Poly
		if err := db.Where("id = ?", id).First(&existingPoly).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Poly not found"})
			return
		}

		// Binding input JSON ke struct Poly
		var updatedPoly model.Poly
		if err := c.ShouldBindJSON(&updatedPoly); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah ID Room yang diberikan ada di database
		var existingRoom model.Room
		if err := db.Where("id = ?", updatedPoly.RoomId).First(&existingRoom).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
			return
		}

		// Update only the specified fields using db.Updates
		if err := db.Model(&existingPoly).Updates(map[string]interface{}{
			"PolyName":        updatedPoly.PolyName,
			"PolyDescription": updatedPoly.PolyDescription,
			"RoomId":          updatedPoly.RoomId,
			"DoctorIdsString": strings.Join(updatedPoly.DoctorIds, ","),
		}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Poly"})
			return
		}

		// Mengambil data Poly yang diperbarui dari database
		var updatedPolyFromDB model.Poly
		if err := db.Where("id = ?", id).First(&updatedPolyFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Poly"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Poly successfully updated",
			Data:    updatedPolyFromDB,
		})
	}
}

func DeletePoly(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var poly model.Poly
		if err := db.Where("id = ?", id).First(&poly).Error; err != nil {
			c.JSON(404, gin.H{"error": "Poly not found"})
			return
		}

		polyName := poly.PolyName

		db.Delete(&poly)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", polyName)})
	}
}

func PermanentDeletePoly(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var poly model.Poly
		if err := db.Unscoped().Where("ID = ?", id).First(&poly).Error; err != nil {
			c.JSON(404, gin.H{"error": "Poly not found"})
			return
		}

		polyName := poly.PolyName

		db.Unscoped().Delete(&poly)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", polyName)})
	}
}
