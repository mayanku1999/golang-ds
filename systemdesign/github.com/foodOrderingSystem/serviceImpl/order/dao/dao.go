package dao

import "github.com/foodOrderingSystem/serviceImpl/order/dao/model"

type IOrderDAO interface {
	CreateOrder(order *model.Order) (*model.Order, error)
	GetOrder(orderID string) (*model.Order, bool)
}
