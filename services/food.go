package services

type DishResponse struct {
	ID   int    `json:"id"`
	Food string `json:"dish"`
}

type DishService interface {
	GetDishes() ([]DishResponse, error)
	GetSingleDish(int) (*DishResponse, error)
	GetDishByKeyword(string) ([]DishResponse, error)
	UpdateDish(int, string) (*DishResponse, error)
}
