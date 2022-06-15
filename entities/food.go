package entities

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

type food struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Food      string `gorm:"foods"`
}

type FoodEntity interface {
	GetAll() ([]food, error)
	GetByID(int) (*food, error)
	GetByKeyword(string) ([]food, error)
	UpdateOne(int, string) (*food, error)
}

func loadData(db *gorm.DB) error {

	// check if data already exist
	var count int64
	db.Model(&food{}).Count(&count)
	if count > 0 {
		return nil
	}

	f, err := os.Open("./datasets/food_dictionary.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	foods := []food{}
	_ = foods

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		foods = append(foods, food{
			Food: fmt.Sprintf(string(line)),
		})
	}

	//fmt.Println(reviews)
	return db.Create(&foods).Error
}

//return db.Create(&reviews).Error
