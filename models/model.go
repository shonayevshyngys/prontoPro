package models

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `json:"username" binding:"required"`
}

type Provider struct {
	ID          uint    `gorm:"primarykey" json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Rating      float32 `json:"rating"`
}

type Review struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	UserId     uint   `json:"userId"  binding:"required"`
	ProviderId uint   `json:"providerId"  binding:"required"`
	ReviewText string `json:"reviewText"`
	Rating     uint8  `json:"rating" binding:"required"`
}
