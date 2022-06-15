package entities

import (
	"gorm.io/gorm"
)

type reviewEntityDB struct {
	db *gorm.DB
}

func NewReviewEntityDB(db *gorm.DB) ReviewEntity {
	db.AutoMigrate(review{})
	loadCSV(db)
	return reviewEntityDB{db: db}
}

func (ent reviewEntityDB) GetAll() ([]review, error) {

	reviews := []review{}
	err := ent.db.Limit(50).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (ent reviewEntityDB) GetByID(id int) (*review, error) {

	review := review{}
	err := ent.db.First(&review, id).Error
	if err != nil {
		return nil, err
	}

	return &review, nil
}

func (ent reviewEntityDB) GetByKeyword(kw string) ([]review, error) {

	reviews := []review{}

	err := ent.db.Limit(100).Where("review LIKE ?", kw).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (ent reviewEntityDB) UpdateOne(id int, update string) (*review, error) {

	review := review{}
	err := ent.db.Model(&review).Where("id = ?", id).Update("review", update).Error
	if err != nil {
		return nil, err
	}

	return &review, nil

}
