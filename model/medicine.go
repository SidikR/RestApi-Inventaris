// model/user.go
package model

import (
	"fmt"
	"main/model/template"
	"main/utils"

	"github.com/jinzhu/gorm"
)

type Medicine struct {
	template.UUIDModel
	MedicineName        string `validate:"required"`
	MedicineDescription string
	Stock               uint   `validate:"required"`
	Unit                string `validate:"required"`
	Image               string
	Expired             string `validate:"required"`
	Price               uint   `validate:"required"`
	template.TimeModel
}

// BeforeCreate callback will be called before data is inserted into the database
func (d *Medicine) BeforeCreate(scope *gorm.Scope) error {
	if err := utils.ValidateStruct(d); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	if err := d.UUIDModel.BeforeCreate(scope); err != nil {
		return fmt.Errorf("error before create: %v", err)
	}
	return nil
}
