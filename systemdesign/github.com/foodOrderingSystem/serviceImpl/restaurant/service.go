package restaurant

import (
	"fmt"
	"github.com/foodOrderingSystem/serviceImpl/restaurant/dao"
	"github.com/foodOrderingSystem/serviceImpl/restaurant/dao/model"
	"sync"
)

var once sync.Once

type RestaurantService struct {
	restaurantDAO dao.IRestaurantDao
}

func NewRestaurantService(restaurantDAO dao.IRestaurantDao) *RestaurantService {
	restaurantService := &RestaurantService{}
	once.Do(func() {
		restaurantService = &RestaurantService{restaurantDAO: restaurantDAO}
	})
	return restaurantService
}

func (service *RestaurantService) OnboardRestaurant(name string, menu map[string]float64, maxCapacity int, currentCap int) (*model.Restaurant, error) {
	if name == "" || len(menu) == 0 || maxCapacity <= 0 {
		return nil, fmt.Errorf("invalid restaurant details")
	}
	res, err := service.restaurantDAO.CreateRestaurant(name, menu, maxCapacity, currentCap)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (service *RestaurantService) UpdateCapacity(restaurantID string, change int) error {
	// TODO put this in transaction block -----------------------------------------------------
	return service.restaurantDAO.UpdateCapacity(restaurantID, change)
}

func (service *RestaurantService) ListAllRestaurants() []*model.Restaurant {
	return service.restaurantDAO.GetAllRestaurants()
}

func (service *RestaurantService) UpdateMenu(restaurantID string, updatedMenu map[string]float64) error {
	return service.restaurantDAO.UpdateMenu(restaurantID, updatedMenu)
}
