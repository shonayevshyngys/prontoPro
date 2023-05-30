package models

import "gorm.io/gorm"

//Rating models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username" binding:"required"`
}

type Provider struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Rating      float32 `json:"rating"`
}

type Review struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	UserID     uint
	User       User `gorm:"foreignKey:user_id" references:"ID"`
	ProviderID uint
	Provider   Provider `gorm:"foreignKey:provider_id" references:"ID"`
	ReviewText string   `json:"reviewText"`
	Rating     uint8    `json:"rating" binding:"required"`
}

//Notification models

type Notification struct {
	gorm.Model   `json:"-"`
	ProviderID   uint
	Notification string
}
