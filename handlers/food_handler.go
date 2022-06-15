package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sun053123/go-csv/services"
)

type dishHandler struct {
	dishServ services.DishService
}

func NewDishHandler(dishServ services.DishService) DishHandler {
	return dishHandler{dishServ: dishServ}
}

func (handl dishHandler) FindDishes(c *fiber.Ctx) error {

	dishes, err := handl.dishServ.GetDishes()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Unexpected Error.")
	}

	response := fiber.Map{
		"status": "ok",
		"dish":   dishes,
	}

	return c.JSON(response)
}

func (handl dishHandler) FindSingleDish(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatal(err)
	}

	dish, err := handl.dishServ.GetSingleDish(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Dish not found.")
	}

	response := fiber.Map{
		"status": "ok",
		"dish":   dish,
	}

	return c.JSON(response)
}

func (handl dishHandler) FindDishByKeyword(c *fiber.Ctx) error {

	kw := c.Params("kw")

	dishes, err := handl.dishServ.GetDishByKeyword(kw)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Dish not found.")
	}

	response := fiber.Map{
		"status": "ok",
		"dish":   dishes,
	}

	return c.JSON(response)
}

func (handl dishHandler) UpdateDish(c *fiber.Ctx) error {

	request := DishRequest{}

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatal(err)
	}

	err = c.BodyParser(&request)
	if err != nil {
		return err
	}

	if request.Update == "" {
		return fiber.ErrUnprocessableEntity
	}

	dish, err := handl.dishServ.UpdateDish(id, request.Update)
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status": "ok",
		"dish":   dish,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
