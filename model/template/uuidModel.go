// uuid_model.go
package template

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UUIDModel struct {
	UUID string `gorm:"type:varchar(36);primary_key;unique_index;not null" validate:"omitempty,uuid4"`
}

func (b *UUIDModel) BeforeCreate(scope *gorm.Scope) error {
	// Generate a new UUID before inserting the data
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	b.UUID = uuid.String()
	return nil
}
