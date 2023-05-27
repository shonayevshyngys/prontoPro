package util

type CreateReviewDTO struct {
	UserId     uint   `json:"userId"  binding:"required"`
	ProviderId uint   `json:"providerId"  binding:"required"`
	ReviewText string `json:"reviewText"`
	Rating     uint8  `json:"rating" binding:"required"`
}
