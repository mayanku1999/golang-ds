package customer

import (
	"fmt"
	"github.com/foodOrderingSystem/serviceImpl/customer/dao"
	"github.com/foodOrderingSystem/serviceImpl/customer/dao/model"
)

type CustomerService struct {
	customerDAO dao.ICustomerDao
}

// NewCustomerService creates and returns a new CustomerService instance
func NewCustomerService(customerDAO dao.ICustomerDao) *CustomerService {
	return &CustomerService{customerDAO: customerDAO}
}

// CreateCustomer creates a new customer with the provided name and stores it in the DAO
func (service *CustomerService) CreateCustomer(name string) *model.Customer {
	return service.customerDAO.CreateCustomer(name)
}

// GetCustomerByID retrieves a customer by its ID from the DAO
func (service *CustomerService) GetCustomerByID(customerID string) (*model.Customer, error) {
	customer, exists := service.customerDAO.GetCustomerByID(customerID)
	if !exists {
		return nil, fmt.Errorf("customer with ID %s not found", customerID)
	}
	return customer, nil
}
