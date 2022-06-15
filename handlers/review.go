package handlers

import "github.com/gofiber/fiber/v2"

type ReviewRequest struct {
	Update string `json:"update"`
}

type ReviewHandler interface {
	FindReviews(c *fiber.Ctx) error
	FindSingleReview(c *fiber.Ctx) error
	FindReviewByKeyword(c *fiber.Ctx) error
	UpdateReview(c *fiber.Ctx) error
}
