package model

import (
	"fmt"
	"main/model/template"
	"main/utils"

	"github.com/jinzhu/gorm"
)

type Doctor struct {
	template.UUIDModel
	DoctorName   string `validate:"required" gorm:"type:varchar(50)"`
	Email        string `validate:"required,email" gorm:"type:varchar(100)"`
	Role         string `validate:"required"`
	JoinAt       string `json:"JoinAt" gorm:"type:varchar(10)"`
	Image        string `json:"Image" gorm:"type:varchar(50)"`
	Status       string `json:"Status" gorm:"type:varchar(50)"`
	MobileNumber string `json:"MobileNumber" gorm:"type:varchar(15)"`
	Address      string `json:"Address"`
	template.TimeModel
}

// BeforeCreate callback will be called before data is inserted into the database
func (d *Doctor) BeforeCreate(scope *gorm.Scope) error {
	if err := utils.ValidateStruct(d); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	if err := d.UUIDModel.BeforeCreate(scope); err != nil {
		return fmt.Errorf("error before create: %v", err)
	}
	return nil
}
