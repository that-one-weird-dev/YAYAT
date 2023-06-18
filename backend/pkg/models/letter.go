package models

import (
	"gorm.io/gorm"
)

type Letter struct {
	gorm.Model
	Text   string
	UserId uint
	User   User
}
