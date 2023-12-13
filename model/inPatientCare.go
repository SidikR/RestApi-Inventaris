package model

import (
	"fmt"
	"main/model/template"
	"main/utils"

	"github.com/jinzhu/gorm"
)

type InPatientCare struct {
	template.UUIDModel
	RegisteredId string       `validate:"required" gorm:"index;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	Registration Registration `gorm:"foreignkey:RegisteredId;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	RoomId       string       `validate:"required" gorm:"index;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	Room         Room         `gorm:"foreignkey:RoomId;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	Days         uint         `validate:"required" gorm:"default:1"`
	Fees         uint
	template.TimeModel
}

// BeforeCreate callback will be called before data is inserted into the database
func (d *InPatientCare) BeforeCreate(scope *gorm.Scope) error {
	if d.Days == 0 {
		d.Days = 1
	}
	if err := utils.ValidateStruct(d); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	if err := d.UUIDModel.BeforeCreate(scope); err != nil {
		return fmt.Errorf("error before create: %v", err)
	}
	return nil
}
