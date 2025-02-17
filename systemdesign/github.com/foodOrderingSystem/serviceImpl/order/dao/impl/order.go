package impl

import (
	"github.com/foodOrderingSystem/serviceImpl/order/dao/model"
	"sync"
)

type OrderDAOImpl struct {
	orders map[string]*model.Order
	mutex  sync.RWMutex
}

func NewOrderDAOImpl() *OrderDAOImpl {
	return &OrderDAOImpl{
		orders: make(map[string]*model.Order),
	}
}

func (dao *OrderDAOImpl) CreateOrder(orderDetails *model.Order) (*model.Order, error) {
	order := model.NewOrder().WithCustomerId(orderDetails.CustomerID).WithTotalCost(orderDetails.TotalCost).WithItemsToRestaurantMap(orderDetails.RestaurantItemAllocation)
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	dao.orders[order.OrderID] = order
	return order, nil
}

func (dao *OrderDAOImpl) GetOrder(orderID string) (*model.Order, bool) {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()
	order, exists := dao.orders[orderID]
	return order, exists
}

func (dao *OrderDAOImpl) GetAllOrders() map[string]*model.Order {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()
	copiedOrders := make(map[string]*model.Order, len(dao.orders))
	for k, v := range dao.orders {
		copiedOrders[k] = v
	}
	return copiedOrders
}
