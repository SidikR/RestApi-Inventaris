package migration

import (
	"main/model"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	// db.AutoMigrate(&model.Doctor{}, &model.InPatientCare{}, &model.Inventory{}, &model.MedicalReport{}, &model.Medicine{}, &model.Patient{}, &model.Poly{}, &model.Receip{}, &model.Registration{}, &model.Room{}, &model.WorkHour{}, &model.User{}, &model.Division{}, &model.BilReport{})
	db.AutoMigrate(&model.MedicalReport{}, &model.Receip{})
}
