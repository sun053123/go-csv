package entities

import (
	"gorm.io/gorm"
)

type foodEntityDB struct {
	db *gorm.DB
}

func NewFoodEntityDB(db *gorm.DB) FoodEntity {
	db.AutoMigrate(food{})
	loadData(db)
	return foodEntityDB{db: db}
}

func (ent foodEntityDB) GetAll() ([]food, error) {

	foods := []food{}
	err := ent.db.Limit(50).Find(&foods).Error
	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (ent foodEntityDB) GetByID(id int) (*food, error) {
	food := food{}
	err := ent.db.First(&food, id).Error
	if err != nil {
		return nil, err
	}

	return &food, nil
}

func (ent foodEntityDB) GetByKeyword(kw string) ([]food, error) {

	foods := []food{}

	err := ent.db.Limit(100).Where("food LIKE ?", kw).Find(&foods).Error
	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (ent foodEntityDB) UpdateOne(id int, update string) (*food, error) {
	food := food{}
	err := ent.db.Model(&food).Where("id = ?", id).Update("food", update).Error
	if err != nil {
		return nil, err
	}

	return &food, nil
}
