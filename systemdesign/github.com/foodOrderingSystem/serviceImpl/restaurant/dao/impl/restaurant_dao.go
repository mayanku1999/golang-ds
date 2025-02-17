package impl

import (
	"fmt"
	"github.com/foodOrderingSystem/pkg/errors"
	"github.com/foodOrderingSystem/serviceImpl/restaurant/dao/model"
	"sync"
)

type RestaurantDAO struct {
	restaurants map[string]*model.Restaurant
	mutex       sync.RWMutex
}

func NewRestaurantDAO() *RestaurantDAO {
	return &RestaurantDAO{restaurants: make(map[string]*model.Restaurant)}
}

func (dao *RestaurantDAO) CreateRestaurant(name string, menu map[string]float64, maxCapacity int, currentCapacity int) (*model.Restaurant, error) {
	if name == "" {
		return nil, errors.InvalidResNameErr
		
	}
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	restaurant := model.NewRestaurant().WithName(name).WithMenu(menu).WithMaxCapacity(maxCapacity).WithCurrentCapacity(currentCapacity)
	dao.restaurants[restaurant.ID] = restaurant
	return restaurant, nil
}

func (dao *RestaurantDAO) GetAllRestaurants() []*model.Restaurant {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	restaurants := make([]*model.Restaurant, 0, len(dao.restaurants))
	for _, r := range dao.restaurants {
		restaurants = append(restaurants, r)
	}
	return restaurants
}

func (dao *RestaurantDAO) UpdateCapacity(restaurantID string, change int) error {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	restaurant, exists := dao.restaurants[restaurantID]
	if !exists {
		return errors.RestaurantNotFoundErr
	}
	newCapacity := restaurant.CurrentCapacity + change
	if newCapacity < 0 || newCapacity > restaurant.MaxCapacity {
		return fmt.Errorf("invalid capacity change, exceeds limits")
	}
	restaurant.CurrentCapacity = newCapacity
	return nil
}

func (dao *RestaurantDAO) GetRestaurantByID(restaurantID string) (*model.Restaurant, bool) {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	restaurant, exists := dao.restaurants[restaurantID]
	return restaurant, exists
}

func (dao *RestaurantDAO) UpdateMenu(restaurantID string, updatedMenu map[string]float64) error {
	if restaurantID == "" || len(updatedMenu) == 0 {
		return errors.InValidMenuErr
	}
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	restaurant, exists := dao.restaurants[restaurantID]
	if !exists {
		return errors.RestaurantNotFoundErr
	}

	for item, price := range updatedMenu {
		if price <= 0 {
			return errors.InvalidMenuPriceErr(item)
		}
		restaurant.Menu[item] = price
	}

	return nil
}
