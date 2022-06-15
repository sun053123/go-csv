package entities

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

type review struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Review    string `gorm:"reviews"`
}

type ReviewEntity interface {
	GetAll() ([]review, error)
	GetByID(int) (*review, error)
	GetByKeyword(string) ([]review, error)
	UpdateOne(int, string) (*review, error)
}

func loadCSV(db *gorm.DB) error {

	// check if data already exist
	var count int64
	db.Model(&review{}).Count(&count)
	if count > 0 {
		return nil
	}

	f, err := os.Open("./datasets/test_file.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(bufio.NewReader(f))
	reader.Comma = ';'
	reader.LazyQuotes = true

	reviews := []review{}
	_ = reviews

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		reviews = append(reviews, review{
			Review: fmt.Sprintf(record[1]),
		})
	}

	//fmt.Println(reviews)
	return db.Create(&reviews).Error

}
