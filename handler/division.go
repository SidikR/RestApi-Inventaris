package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateDivision(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Division
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the Division to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Division"})
			return
		}

		// Respond with HTTP 201 Created status and the created Division data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Division created succesfully",
			Data:    input,
		})
	}
}

func GetDivision(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.Division

		// Mengambil data dari tabel "division" dengan join ke tabel "division"
		if err := db.Table("divisions").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Not Found Division"})
			return
		}

		c.JSON(200, result)
	}
}

func GetDivisionByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var division model.Division
		if err := db.Where("id = ?", id).First(&division).Error; err != nil {
			c.JSON(404, gin.H{"error": "Division not found"})
			return
		}
		c.JSON(200, division)
	}
}

func UpdateDivision(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		id := c.Param("id")

		// Mencari division berdasarkan UUID
		var existingDivision model.Division
		if err := db.Where("id = ?", id).First(&existingDivision).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Division not found"})
			return
		}

		// Binding input JSON ke struct Division
		var updatedDivision model.Division
		if err := c.ShouldBindJSON(&updatedDivision); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Memperbarui atribut Division
		db.Model(&existingDivision).Updates(&updatedDivision)

		// Mengambil data Division yang diperbarui dari database
		var updatedDivisionFromDB model.Division
		if err := db.Where("id = ?", id).First(&updatedDivisionFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Division"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Division successfully updated",
			Data:    updatedDivisionFromDB,
		})

	}
}

func DeleteDivision(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var division model.Division
		if err := db.Where("id = ?", id).First(&division).Error; err != nil {
			c.JSON(404, gin.H{"error": "Division not found"})
			return
		}

		divisionName := division.DivisionName

		db.Delete(&division)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", divisionName)})
	}
}

func PermanentDeleteDivision(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var division model.Division
		if err := db.Unscoped().Where("ID = ?", id).First(&division).Error; err != nil {
			c.JSON(404, gin.H{"error": "Division not found"})
			return
		}

		divisionName := division.DivisionName

		db.Unscoped().Delete(&division)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", divisionName)})
	}
}
