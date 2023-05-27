package models

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
	User       User `gorm:"foreignKey:ID" references:"ID"`
	ProviderID uint
	Provider   Provider `gorm:"foreignKey:ID" references:"ID"`
	ReviewText string   `json:"reviewText"`
	Rating     uint8    `json:"rating" binding:"required"`
}
