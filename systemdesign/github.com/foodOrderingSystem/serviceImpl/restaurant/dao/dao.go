package dao

import (
	"github.com/foodOrderingSystem/serviceImpl/restaurant/dao/model"
)

type IRestaurantDao interface {
	CreateRestaurant(name string, menu map[string]float64, maxCapacity int, currentCap int) (*model.Restaurant, error)
	GetAllRestaurants() []*model.Restaurant
	UpdateCapacity(restaurantID string, change int) error
	GetRestaurantByID(restaurantID string) (*model.Restaurant, bool)
	UpdateMenu(restaurantID string, updatedMenu map[string]float64) error
}
