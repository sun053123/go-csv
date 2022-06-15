package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sun053123/go-csv/entities"
	"github.com/sun053123/go-csv/handlers"
	"github.com/sun053123/go-csv/services"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB) {

	foodEntity := entities.NewFoodEntityDB(db)
	foodService := services.NewDishService(foodEntity)
	foodHandler := handlers.NewDishHandler(foodService)

	app.Get("/dishes", foodHandler.FindDishes)
	app.Get("/dishes:kw?", foodHandler.FindDishByKeyword)
	app.Get("/dishes/:id", foodHandler.FindSingleDish)
	app.Post("/dishes/:id", foodHandler.UpdateDish)

	reviewEntity := entities.NewReviewEntityDB(db)
	reviewService := services.NewReviewService(reviewEntity)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	app.Get("/reviews", reviewHandler.FindReviews)
	app.Get("/reviews:kw?", reviewHandler.FindReviewByKeyword)
	app.Get("/reviews/:id", reviewHandler.FindSingleReview)
	app.Post("/reviews/:id", reviewHandler.UpdateReview)
}
