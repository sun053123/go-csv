package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sun053123/go-csv/services"
)

type reviewHandler struct {
	reviewServ services.ReviewService
}

func NewReviewHandler(reviewServ services.ReviewService) ReviewHandler {
	return reviewHandler{reviewServ: reviewServ}
}

func (handl reviewHandler) FindReviews(c *fiber.Ctx) error {

	reivews, err := handl.reviewServ.GetReviews()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unexpected Error")
	}

	response := fiber.Map{
		"status":  fiber.StatusOK,
		"reviews": reivews,
	}

	return c.JSON(response)

}

func (handl reviewHandler) FindSingleReview(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unexpected Error")
	}

	review, err := handl.reviewServ.GetSingleReview(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Not found.")
	}

	response := fiber.Map{
		"status":  fiber.StatusOK,
		"reviews": review,
	}

	return c.JSON(response)
}

func (handl reviewHandler) FindReviewByKeyword(c *fiber.Ctx) error {

	kw := c.Params("kw")

	reviews, err := handl.reviewServ.GetReviewByKeyword(kw)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	response := fiber.Map{
		"status":  fiber.StatusOK,
		"reviews": reviews,
	}

	return c.JSON(response)
}

func (handl reviewHandler) UpdateReview(c *fiber.Ctx) error {

	request := ReviewRequest{}

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unexpected Error")
	}

	err = c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = c.BodyParser(&request)
	if err != nil {
		return err
	}

	if request.Update == "" {
		return fiber.ErrUnprocessableEntity
	}

	review, err := handl.reviewServ.UpdateReview(id, request.Update)
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status": fiber.StatusCreated,
		"reivew": review,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
