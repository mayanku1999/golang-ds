package model

import (
	"fmt"
	"math/rand"
)

type Order struct {
	OrderID                  string
	CustomerID               string
	TotalCost                float64
	RestaurantItemAllocation map[string]map[string]int
}

func NewOrder() *Order {
	return &Order{
		OrderID: fmt.Sprintf("ORD-%d", rand.Intn(1000000)),
	}
}

func (o *Order) WithCustomerId(customerId string) *Order {
	o.CustomerID = customerId
	return o
}
func (o *Order) WithTotalCost(totalCost float64) *Order {
	o.TotalCost = totalCost
	return o
}
func (o *Order) WithItemsToRestaurantMap(itemsToRestaurantMap map[string]map[string]int) *Order {
	o.RestaurantItemAllocation = itemsToRestaurantMap
	return o
}
