package impl

import (
	"github.com/foodOrderingSystem/serviceImpl/customer/dao/model"
	"math/rand"
	"strconv"
	"sync"
)

var syncOnce sync.Once

type CustomerDAO struct {
	customers map[string]*model.Customer
	mutex     sync.RWMutex
}

func NewCustomerDAO() *CustomerDAO {
	var customerSvc *CustomerDAO
	syncOnce.Do(func() {
		customerSvc = &CustomerDAO{
			customers: make(map[string]*model.Customer),
		}
	})
	return customerSvc
}

func (dao *CustomerDAO) CreateCustomer(name string) *model.Customer {
	customerID := "CUST-" + strconv.Itoa(rand.Intn(1000000))
	customer := &model.Customer{
		ID:   customerID,
		Name: name,
	}
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	dao.customers[customerID] = customer

	return customer
}

func (dao *CustomerDAO) GetCustomerByID(customerID string) (*model.Customer, bool) {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()
	customer, exists := dao.customers[customerID]
	return customer, exists
}
