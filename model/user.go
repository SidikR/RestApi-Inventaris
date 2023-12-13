// model/user.go
package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique_index" validate:"required,email"`
	Username string `gorm:"unique_index" validate:"required"`
	Password string `validate:"required"`
	Role     string `validate:"required"`
	Image    string
	Token    string
}
