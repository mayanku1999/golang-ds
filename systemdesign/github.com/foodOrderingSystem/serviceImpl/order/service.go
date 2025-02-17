package order

import (
	"fmt"
	"github.com/foodOrderingSystem/pkg/errors"
	"github.com/foodOrderingSystem/serviceImpl/customer"
	"github.com/foodOrderingSystem/serviceImpl/order/dao"
	"github.com/foodOrderingSystem/serviceImpl/order/dao/model"
	"github.com/foodOrderingSystem/serviceImpl/restaurant"
	"github.com/foodOrderingSystem/serviceImpl/restaurant/service/selection_strategy"
)

type OrderService struct {
	orderDAO          dao.IOrderDAO
	restaurantService *restaurant.RestaurantService
	customerService   *customer.CustomerService
}

func NewOrderSvc(orderDAO dao.IOrderDAO, restaurantService *restaurant.RestaurantService, customerService *customer.CustomerService) *OrderService {
	return &OrderService{
		orderDAO:          orderDAO,
		restaurantService: restaurantService,
		customerService:   customerService,
	}
}

func (service *OrderService) PlaceOrder(customerID string, items map[string]int, strategyType selection_strategy.Strategy) (*model.Order, error) {
	// Check if the customer exists
	_, err := service.customerService.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}

	// Determine the strategy to use
	strategyInstance, err := selection_strategy.GetRestaurantSelectionStrategy(strategyType, service.restaurantService)
	if err != nil {
		return nil, err
	}
	
	// Use the strategy to select restaurants for the items
	restaurantItemAllocation, totalCost, err := strategyInstance.SelectRestaurant(items)
	if restaurantItemAllocation == nil || err != nil {
		return nil, fmt.Errorf("unable to fulfill the order with the available restaurants: %v", err)
	}

	restaurantItemMap := make(map[string]map[string]int)

	// Iterate over the allocations
	for _, allocation := range restaurantItemAllocation {
		if _, exists := restaurantItemMap[allocation.RestaurantId]; !exists {
			restaurantItemMap[allocation.RestaurantId] = make(map[string]int)
		}

		if _, itemExists := restaurantItemMap[allocation.RestaurantId][allocation.Item]; !itemExists {
			restaurantItemMap[allocation.RestaurantId][allocation.Item] = 0
		}

		restaurantItemMap[allocation.RestaurantId][allocation.Item] += allocation.Quantity
		err := service.restaurantService.UpdateCapacity(allocation.RestaurantId, allocation.Quantity)
		if err != nil {
			return nil, errors.UpdatingCapErr(allocation.RestaurantId, err)
		}
	}

	order := &model.Order{
		CustomerID:               customerID,
		TotalCost:                totalCost,
		RestaurantItemAllocation: restaurantItemMap,
	}
	orderRes, err := service.orderDAO.CreateOrder(order)
	if err != nil {
		return nil, errors.OrderCreationFailureErr(err)
	}
	return orderRes, nil
}

func (service *OrderService) FulfillOrder(orderID string) error {
	order, exists := service.orderDAO.GetOrder(orderID)
	if !exists {
		return errors.OrderNotFound
	}

	for resId, item := range order.RestaurantItemAllocation {
		for _, qty := range item {
			err := service.restaurantService.UpdateCapacity(resId, -qty)
			if err != nil {
				return errors.UpdatingCapErr(resId, err)
			}
		}
	}
	return nil
}
