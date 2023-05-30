package util

type ErrorMessage struct {
	Code    uint
	Message string
}
type SuccessMessage struct {
	Code    uint
	Message string
}

type Subscribers struct {
	Ids map[int]string
}

type CreateReviewDTO struct {
	UserId     uint   `json:"userId"  binding:"required"`
	ProviderId uint   `json:"providerId"  binding:"required"`
	ReviewText string `json:"reviewText"`
	Rating     uint8  `json:"rating" binding:"required"`
}

type SubscriptionBody struct {
	UserId     int
	ProviderId int
}

const BadIdText = "Bad format for id"
