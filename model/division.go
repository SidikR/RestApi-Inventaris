package model

import (
	"github.com/jinzhu/gorm"
)

type Division struct {
	gorm.Model
	DivisionName        string `validate:"required"`
	DivisionDescription string
}
