package selection_strategy

import (
	"github.com/foodOrderingSystem/pkg/errors"
	"github.com/foodOrderingSystem/serviceImpl/restaurant"
)

type IRestaurantSelectionStrategy interface {
	SelectRestaurant(items map[string]int) ([]Allocation, float64, error)
}

func GetRestaurantSelectionStrategy(strategyType Strategy, service *restaurant.RestaurantService) (IRestaurantSelectionStrategy, error) {
	var strategyInstance IRestaurantSelectionStrategy
	switch strategyType {
	case LOWEST_PRICE_STRATEGY:
		strategyInstance = NewLowestPriceStrategy(service)
	default:
		return nil, errors.UndefinedSelectionStrategy
	}
	return strategyInstance, nil
}
