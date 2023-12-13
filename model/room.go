// model/user.go
package model

import (
	"github.com/jinzhu/gorm"
)

type Room struct {
	gorm.Model
	RoomName        string `validate:"required"`
	RoomDescription string
	Location        string `validate:"required"`
	DivisionId      uint   `validate:"required" gorm:"index;foreignkey:DivisionId;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	// Division        Division `gorm:"foreignkey:DivisionId;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
}
