package handlers

import "github.com/gofiber/fiber/v2"

type DishRequest struct {
	ID     string `json:"id"`
	Update string `json:"update_dish"`
}

type DishHandler interface {
	FindDishes(c *fiber.Ctx) error
	FindSingleDish(c *fiber.Ctx) error
	FindDishByKeyword(c *fiber.Ctx) error
	UpdateDish(c *fiber.Ctx) error
}
