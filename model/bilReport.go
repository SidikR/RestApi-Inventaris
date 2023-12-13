package model

import (
	"fmt"
	"main/model/template"
	"main/utils"

	"github.com/jinzhu/gorm"
)

type BilReport struct {
	template.UUIDModel
	RegisteredId string
	RecipeId     string
	Recipe       Receip `gorm:"foreignKey:RecipeId"`
	Services     uint
	Total        uint
	template.TimeModel
}

// BeforeCreate callback will be called before data is inserted into the database
func (d *BilReport) BeforeCreate(scope *gorm.Scope) error {
	if err := utils.ValidateStruct(d); err != nil {
		return fmt.Errorf("validation error detected: %v", err)
	}

	if err := d.UUIDModel.BeforeCreate(scope); err != nil {
		return fmt.Errorf("error before create: %v", err)
	}

	return nil
}
