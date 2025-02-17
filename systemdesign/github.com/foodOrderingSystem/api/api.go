package api

import (
	"github.com/foodOrderingSystem/serviceImpl/order/dao/model"
	resModel "github.com/foodOrderingSystem/serviceImpl/restaurant/dao/model"
	"github.com/foodOrderingSystem/serviceImpl/restaurant/service/selection_strategy"
)

type IRestaurantService interface {
	OnboardRestaurant(name string, menu map[string]float64, maxCapacity int, currentCap int) (*resModel.Restaurant, error)
	UpdateCapacity(restaurantID string, change int) error
	ListAllRestaurants() []*resModel.Restaurant
	UpdateMenu(restaurantID string, updatedMenu map[string]float64) error
}

type IOrderService interface {
	PlaceOrder(customerID string, items map[string]int, strategyType selection_strategy.Strategy) (*model.Order, error)
	FulfillOrder(orderID string) error
}
