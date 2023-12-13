package model

import (
	"fmt"
	"main/model/template"
	"main/utils"

	"github.com/jinzhu/gorm"
)

type MedicalReport struct {
	template.UUIDModel
	RegisteredId string
	// RecipeId     string
	BilId string
	template.TimeModel
}

// BeforeCreate callback will be called before data is inserted into the database
func (d *MedicalReport) BeforeCreate(scope *gorm.Scope) error {
	if err := utils.ValidateStruct(d); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	if err := d.UUIDModel.BeforeCreate(scope); err != nil {
		return fmt.Errorf("error before create: %v", err)
	}
	return nil
}
