package services

import (
	"fmt"

	"github.com/sun053123/go-csv/entities"
)

type reviewService struct {
	reviewEnt entities.ReviewEntity
}

func NewReviewService(reviewEnt entities.ReviewEntity) ReviewService {
	return reviewService{reviewEnt: reviewEnt}
}

func (serv reviewService) GetReviews() ([]ReviewResponse, error) {

	reviews := []ReviewResponse{}

	reviewDB, err := serv.reviewEnt.GetAll()
	if err != nil {
		return nil, err
	}

	for _, review := range reviewDB {
		reviews = append(reviews, ReviewResponse{
			ID:     review.ID,
			Review: review.Review,
		})
	}

	return reviews, nil
}

func (serv reviewService) GetSingleReview(id int) (*ReviewResponse, error) {

	reviewDB, err := serv.reviewEnt.GetByID(id)
	if err != nil {
		return nil, err
	}

	review := ReviewResponse{
		ID:     reviewDB.ID,
		Review: reviewDB.Review,
	}

	return &review, nil
}

func (serv reviewService) GetReviewByKeyword(kw string) ([]ReviewResponse, error) {

	reviews := []ReviewResponse{}

	keyword := fmt.Sprintf("%%%v%%", kw)

	reviewDB, err := serv.reviewEnt.GetByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	for _, review := range reviewDB {
		reviews = append(reviews, ReviewResponse{
			ID:     review.ID,
			Review: review.Review,
		})
	}

	return reviews, nil
}

func (serv reviewService) UpdateReview(id int, update string) (*ReviewResponse, error) {

	reviewDB, err := serv.reviewEnt.UpdateOne(id, update)
	if err != nil {
		return nil, err
	}

	review := ReviewResponse{
		ID:     reviewDB.ID,
		Review: reviewDB.Review,
	}

	return &review, nil
}
