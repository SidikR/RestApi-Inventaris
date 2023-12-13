package model

import (
	"fmt"
	"main/model/template"
	"main/utils"

	"github.com/jinzhu/gorm"
)

type Receip struct {
	template.UUIDModel
	RegisteredId  string `validate:"required" gorm:"index;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	Diagnosis     string
	ListMedicines []ReceiptMedicine `gorm:"many2many:receipt_medicines"`
	TotalPrice    uint
	template.TimeModel
}

type ReceiptMedicine struct {
	ReceipUUID  string `gorm:"index;constraint:OnDelete:CASCADE;constraint:OnUpdate:CASCADE"`
	MedicineID  string
	Qty         uint
	Dose        string
	Description string
}

// BeforeCreate callback will be called before data is inserted into the database
func (r *Receip) BeforeCreate(scope *gorm.Scope) error {
	if err := utils.ValidateStruct(r); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	// Generate a new UUID before inserting the data
	uuidString, err := utils.GenerateUUID()
	if err != nil {
		return fmt.Errorf("error generating UUID: %v", err)
	}

	// Set the UUID in the model using scope.SetColumn
	scope.SetColumn("UUID", uuidString)

	// Set the ReceipUUID in related ReceiptMedicine instances
	for i := range r.ListMedicines {
		r.ListMedicines[i].ReceipUUID = uuidString
	}

	return nil
}
