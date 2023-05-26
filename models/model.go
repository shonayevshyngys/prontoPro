package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Username   string `json:"username" binding:"required"`
}

type Provider struct {
	gorm.Model  `json:"-"`
	Description string  `json:"description"`
	Rating      float32 `json:":rating"`
}

type Review struct {
	gorm.Model `json:"-"`
	UserId     uint   `json:"userId"`
	ProviderId uint   `json:"providerId"`
	ReviewText string `json:"reviewText"`
}
