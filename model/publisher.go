package model

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name    string
	Address string
	Books   []Book `gorm:"foreignKey:PublisherID"`
}
