package services

import (
	"fmt"

	"github.com/sun053123/go-csv/entities"
)

type dishService struct {
	foodEnt entities.FoodEntity
}

func NewDishService(foodEnt entities.FoodEntity) DishService {
	return dishService{foodEnt: foodEnt}
}

func (serv dishService) GetDishes() ([]DishResponse, error) {

	dishes := []DishResponse{}

	foodDB, err := serv.foodEnt.GetAll()
	if err != nil {
		return nil, err
	}

	for _, dish := range foodDB {
		dishes = append(dishes, DishResponse{
			ID:   dish.ID,
			Food: dish.Food,
		})
	}

	return dishes, nil
}

func (serv dishService) GetSingleDish(id int) (*DishResponse, error) {

	food, err := serv.foodEnt.GetByID(id)
	if err != nil {
		return nil, err
	}

	dish := DishResponse{
		ID:   food.ID,
		Food: food.Food,
	}

	return &dish, nil
}

func (serv dishService) GetDishByKeyword(kw string) ([]DishResponse, error) {

	dishes := []DishResponse{}

	keyword := fmt.Sprintf("%%%v%%", kw)

	foodDB, err := serv.foodEnt.GetByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	for _, dish := range foodDB {
		dishes = append(dishes, DishResponse{
			ID:   dish.ID,
			Food: dish.Food,
		})
	}

	return dishes, nil
}

func (serv dishService) UpdateDish(id int, update string) (*DishResponse, error) {

	food, err := serv.foodEnt.UpdateOne(id, update)
	if err != nil {
		return nil, err
	}

	dish := DishResponse{
		ID:   food.ID,
		Food: food.Food,
	}

	return &dish, nil
}
