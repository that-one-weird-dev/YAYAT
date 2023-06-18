package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DisplayName string
	Tag         string `gorm:"uniqueIndex"`
	Password    string
	Letters     []Letter
}
