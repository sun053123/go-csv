package services

type ReviewResponse struct {
	ID     int    `json:"id"`
	Review string `json:"review"`
}

type ReviewService interface {
	GetReviews() ([]ReviewResponse, error)
	GetSingleReview(int) (*ReviewResponse, error)
	GetReviewByKeyword(string) ([]ReviewResponse, error)
	UpdateReview(int, string) (*ReviewResponse, error)
}
