package model

import (
	"fmt"
	"main/model/template"
	"main/utils"

	"github.com/jinzhu/gorm"
)

type Registration struct {
	template.UUIDModel
	PatientId string  `validate:"required" gorm:"index;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	Patient   Patient `gorm:"foreignKey:PatientId"`
	DoctorId  string  `gorm:"index;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	PolyId    uint    `validate:"required" gorm:"index;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	Status    string
	template.TimeModel
}

// BeforeCreate callback will be called before data is inserted into the database
func (d *Registration) BeforeCreate(scope *gorm.Scope) error {
	if err := utils.ValidateStruct(d); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	if err := d.UUIDModel.BeforeCreate(scope); err != nil {
		return fmt.Errorf("error before create: %v", err)
	}
	return nil
}
