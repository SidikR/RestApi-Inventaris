package model

import (
	"github.com/jinzhu/gorm"
)

type Inventory struct {
	gorm.Model
	InventoryName        string `validate:"required"`
	InventoryDescription string
	Qty                  uint `validate:"required"`
	Condition            string
	RoomId               uint `validate:"required" gorm:"foreignkey:RoomId;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
}
