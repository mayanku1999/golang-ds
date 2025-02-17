package dao

import "github.com/foodOrderingSystem/serviceImpl/customer/dao/model"

type ICustomerDao interface {
	CreateCustomer(name string) *model.Customer
	GetCustomerByID(customerID string) (*model.Customer, bool)
}
