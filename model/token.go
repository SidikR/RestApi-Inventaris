// model/token.go
package model

import "github.com/jinzhu/gorm"

type Token struct {
	gorm.Model
	UserID uint
	Value  string
	Expiry int64
}
