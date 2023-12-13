// model/poly.go
package model

import (
	"fmt"
	"main/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

type Poly struct {
	gorm.Model
	PolyName        string `validate:"required"`
	PolyDescription string
	RoomId          uint     `gorm:"index;constraint:OnDelete:NO ACTION;constraint:OnUpdate:CASCADE"`
	DoctorIds       []string `gorm:"-"`
	DoctorIdsString string
}

// BeforeCreate callback will be called before data is inserted into the database
func (p *Poly) BeforeCreate(scope *gorm.Scope) error {
	if err := utils.ValidateStruct(p); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	// Convert DoctorIds to a string with a separator (comma)
	p.DoctorIdsString = strings.Join(p.DoctorIds, ",")

	return nil
}
