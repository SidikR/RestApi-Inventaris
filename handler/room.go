// handler/room_handler.go
package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateRoom(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Room
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the provided IDDepartemen exists in the database
		var DivisionId model.Division
		if err := db.Where("id = ?", input.DivisionId).First(&DivisionId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Divisi not found"})
			return
		}

		// Save the Room to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Room"})
			return
		}

		// Respond with HTTP 201 Created status and the created Room data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Room berhasil dibuat",
			Data:    input,
		})
	}
}

func GetRoom(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []struct {
			model.Room
			DivisionName string
		}

		db.LogMode(true)
		if err := db.Table("rooms").
			Select("rooms.*, divisions.division_name").
			Joins("JOIN divisions ON rooms.division_id = divisions.id").
			Scan(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": "Gagal mengambil data room"})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data room"})
			return
		}

		c.JSON(200, result)
	}
}

func GetRoomByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var room struct {
			model.Room
			DivisionName string
		}

		// Aktifkan logging untuk melihat pernyataan SQL
		db.LogMode(true)

		// Mengambil data dari tabel "room" dengan join ke tabel "division" berdasarkan ID
		if err := db.Table("rooms").
			Select("rooms.*, divisions.division_name").
			Joins("JOIN divisions ON rooms.division_id = divisions.id").
			Where("rooms.ID = ?", id).
			First(&room).Error; err != nil {
			c.JSON(404, gin.H{"error": "Room not found"})
			return
		}

		c.JSON(200, room)
	}
}

func UpdateRoom(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan ID dari URL parameter
		id := c.Param("id")

		// Mencari room berdasarkan ID
		var existingRoom model.Room
		if err := db.Where("id = ?", id).First(&existingRoom).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
			return
		}

		// Binding input JSON ke struct Room
		var updatedRoom model.Room
		if err := c.ShouldBindJSON(&updatedRoom); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah IDDepartemen yang diberikan ada di database
		var existingDepartemen model.Division
		if err := db.Where("id = ?", updatedRoom.DivisionId).First(&existingDepartemen).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Division not found"})
			return
		}

		// Memperbarui atribut Room
		db.Model(&existingRoom).Updates(&updatedRoom)

		// Mengambil data Room yang diperbarui dari database
		var updatedRoomFromDB model.Room
		if err := db.Where("id = ?", id).First(&updatedRoomFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Room"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Room successfully updated",
			Data:    updatedRoomFromDB,
		})

	}
}

func DeleteRoom(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var room model.Room
		if err := db.Where("id = ?", id).First(&room).Error; err != nil {
			c.JSON(404, gin.H{"error": "Room not found"})
			return
		}

		roomName := room.RoomName

		db.Delete(&room)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", roomName)})
	}
}

func PermanentDeleteRoom(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var room model.Room
		if err := db.Unscoped().Where("ID = ?", id).First(&room).Error; err != nil {
			c.JSON(404, gin.H{"error": "Room not found"})
			return
		}

		roomName := room.RoomName

		db.Unscoped().Delete(&room)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", roomName)})
	}
}
