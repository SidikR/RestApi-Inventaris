// handler/inventory_handler.go
package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateInventory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Inventory
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

		// Save the Inventory to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Inventory"})
			return
		}

		// Respond with HTTP 201 Created status and the created Inventory data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Inventory berhasil dibuat",
			Data:    input,
		})
	}
}

func GetInventory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []struct {
			model.Inventory
			RoomName string
		}

		// Mengambil data dari tabel "inventories" dengan join ke tabel "rooms"
		if err := db.Table("inventories").
			Select("inventories.* , rooms.room_name").
			Joins("JOIN rooms ON inventories.room_id = rooms.id").
			Scan(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data inventory"})
			return
		}

		c.JSON(200, result)
	}
}

func GetInventoryByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var inventory model.Inventory
		if err := db.Where("id = ?", id).First(&inventory).Error; err != nil {
			c.JSON(404, gin.H{"error": "Inventory not found"})
			return
		}
		c.JSON(200, inventory)
	}
}

func UpdateInventory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan ID dari URL parameter
		id := c.Param("id")

		// Mencari inventory berdasarkan ID
		var existingInventory model.Inventory
		if err := db.Where("id = ?", id).First(&existingInventory).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
			return
		}

		// Binding input JSON ke struct Inventory
		var updatedInventory model.Inventory
		if err := c.ShouldBindJSON(&updatedInventory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah IDRoom yang diberikan ada di database
		var existingRoom model.Room
		if err := db.Where("id = ?", updatedInventory.RoomId).First(&existingRoom).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
			return
		}

		// Memperbarui atribut Inventory
		db.Model(&existingInventory).Updates(&updatedInventory)

		// Mengambil data Inventory yang diperbarui dari database
		var updatedInventoryFromDB model.Inventory
		if err := db.Where("id = ?", id).First(&updatedInventoryFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Inventory"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Inventory successfully updated",
			Data:    updatedInventoryFromDB,
		})

	}
}

func DeleteInventory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var inventory model.Inventory
		if err := db.Where("id = ?", id).First(&inventory).Error; err != nil {
			c.JSON(404, gin.H{"error": "Inventory not found"})
			return
		}

		inventoryName := inventory.InventoryName

		db.Delete(&inventory)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", inventoryName)})
	}
}

func PermanentDeleteInventory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var inventory model.Inventory
		if err := db.Unscoped().Where("ID = ?", id).First(&inventory).Error; err != nil {
			c.JSON(404, gin.H{"error": "Inventory not found"})
			return
		}

		inventoryName := inventory.InventoryName

		db.Unscoped().Delete(&inventory)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", inventoryName)})
	}
}
