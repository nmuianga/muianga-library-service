package model

import "gorm.io/gorm"

type BookCategory struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Books []Book `gorm:"foreignKey:CategoryID"`
}
