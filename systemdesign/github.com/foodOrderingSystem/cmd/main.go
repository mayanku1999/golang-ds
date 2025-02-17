package main

import (
	"fmt"
	"time"

	"github.com/foodOrderingSystem/serviceImpl/customer"
	impl2 "github.com/foodOrderingSystem/serviceImpl/customer/dao/impl"
	order2 "github.com/foodOrderingSystem/serviceImpl/order"
	impl3 "github.com/foodOrderingSystem/serviceImpl/order/dao/impl"
	"github.com/foodOrderingSystem/serviceImpl/order/dao/model"
	"github.com/foodOrderingSystem/serviceImpl/restaurant"
	"github.com/foodOrderingSystem/serviceImpl/restaurant/dao/impl"
	"github.com/foodOrderingSystem/serviceImpl/restaurant/service/selection_strategy"
)

func main() {
	// Initialize DAOs
	customerDAO := impl2.NewCustomerDAO()
	restaurantDAO := impl.NewRestaurantDAO()
	orderDAO := impl3.NewOrderDAOImpl()

	// Initialize Services
	restaurantSvcClient := restaurant.NewRestaurantService(restaurantDAO)
	customerSvcClient := customer.NewCustomerService(customerDAO)
	orderServiceClient := order2.NewOrderSvc(orderDAO, restaurantSvcClient, customerSvcClient)

	// Add Restaurants
	rest1, err := restaurantSvcClient.OnboardRestaurant(
		"resta1",
		map[string]float64{"burger": 10, "pizza": 120},
		10, 0,
	)

	if err != nil {
		fmt.Println("Failed to onboard resta1:", err)
		return
	}

	if _, err := restaurantSvcClient.OnboardRestaurant(
		"resta2",
		map[string]float64{"burger": 8, "pizza": 1},
		12, 0,
	); err != nil {
		fmt.Println("Failed to onboard resta2:", err)
		return
	}

	if _, err := restaurantSvcClient.OnboardRestaurant(
		"resta3",
		map[string]float64{"burger": 12},
		8, 0,
	); err != nil {
		fmt.Println("Failed to onboard resta2:", err)
		return
	}

	//Update Menu
	if err := restaurantSvcClient.UpdateMenu(rest1.ID, map[string]float64{
		"bendi_macaroni": 8,
		"king_burger":    15,
	}); err != nil {
		fmt.Println("Failed to update menu for resta1:", err)
		return
	}

	// Add a Customer
	customer := customerSvcClient.CreateCustomer("cust1")

	// Place an Order
	orderItems := map[string]int{
		"burger": 5,
		"pizza":  3,
	}
	fmt.Println("order items", orderItems)

	// Print all restaurant details
	restaurants := restaurantSvcClient.ListAllRestaurants()
	fmt.Println("All restaurants:")
	for _, restaurant := range restaurants {
		fmt.Printf("%+v\n", restaurant)
	}

	order, err := orderServiceClient.PlaceOrder(customer.ID, orderItems, selection_strategy.LOWEST_PRICE_STRATEGY)
	if err != nil {
		fmt.Println("Order placement failed:", err)
		return
	}
	fmt.Printf("Order placed successfully: %+v\n", order)

	// Print all restaurant details after order placement
	restaurants = restaurantSvcClient.ListAllRestaurants()
	fmt.Println("All restaurants after order placement:")
	for _, restaurant := range restaurants {
		fmt.Printf("%+v\n", restaurant)
	}

	// Simulate order fulfillment
	go func(o *model.Order) {
		time.Sleep(5 * time.Second) // Simulate preparation time
		if err := orderServiceClient.FulfillOrder(o.OrderID); err != nil {
			fmt.Println("Failed to fulfill order:", err)
		} else {
			fmt.Printf("Order fulfilled successfully: %+v\n", o)
		}
	}(order)

	// Print all orders
	orders := orderDAO.GetAllOrders()
	fmt.Println("All orders:")
	for _, o := range orders {
		fmt.Printf("%+v\n", o)
	}

	// Keep the main thread alive
	time.Sleep(10 * time.Second)

	// Print all restaurant details after order placement
	restaurants = restaurantSvcClient.ListAllRestaurants()
	fmt.Println("All restaurants after order fulfillment:")
	for _, restaurant := range restaurants {
		fmt.Printf("%+v\n", restaurant)
	}
}
