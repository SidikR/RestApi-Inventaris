package model

import "github.com/jinzhu/gorm"

type WorkHour struct {
	gorm.Model
	DoctorId string `validate:"required" gorm:"foreignkey:DoctorId;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	StartAt  string `validate:"required"`
	FinishAt string `validate:"required"`
}
