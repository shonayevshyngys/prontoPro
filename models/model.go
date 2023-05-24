package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
}

type Provider struct {
	gorm.Model
	Description string
	Rating float32
}