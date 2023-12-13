package model

import (
	"main/model/template"
)

type Patient struct {
	NIK          string `gorm:"primary_key;unique_index;index" validate:"required"`
	PatientName  string `validate:"required"`
	BpjsNumber   string `validate:"required"`
	MobileNumber string `validate:"required"`
	Address      string `validate:"required"`
	template.TimeModel
}
