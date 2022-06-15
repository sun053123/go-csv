package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sun053123/go-csv/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sun053123/go-csv/routes"
)

func main() {

	PORT := os.Getenv("APP_PORT")

	db := initDB()

	reviewEntity := entities.NewReviewEntityDB(db)
	_ = reviewEntity

	app := fiber.New(fiber.Config{
		//Prefork: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("api ready!")
	})

	routes.Routes(app, db)

	fmt.Printf("server ready at http://localhost%s", PORT)
	app.Listen(PORT)

}

func initDB() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", host, user, password, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
