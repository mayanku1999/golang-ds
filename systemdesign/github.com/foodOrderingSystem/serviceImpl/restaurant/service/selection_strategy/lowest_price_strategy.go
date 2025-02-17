package selection_strategy

import (
	"fmt"
	"github.com/foodOrderingSystem/serviceImpl/restaurant"
	"github.com/foodOrderingSystem/serviceImpl/restaurant/dao/model"
	"math"
)

type LowestPriceStrategy struct {
	restaurantSvcClient *restaurant.RestaurantService
}

type Allocation struct {
	RestaurantId string
	Item         string
	Quantity     int
	Cost         float64
}

func NewLowestPriceStrategy(restaurantSvc *restaurant.RestaurantService) *LowestPriceStrategy {
	return &LowestPriceStrategy{
		restaurantSvcClient: restaurantSvc,
	}
}

func (s *LowestPriceStrategy) SelectRestaurant(items map[string]int) ([]Allocation, float64, error) {
	restaurants := s.restaurantSvcClient.ListAllRestaurants()
	return findMinimumCostWithAllocation(restaurants, items)
}

func findMinimumCostWithAllocation(restaurants []*model.Restaurant, order map[string]int) ([]Allocation, float64, error) {
	var optimalAllocations []Allocation
	minCost := math.MaxFloat64
	successful := false

	var backtrack func(itemIndex int, allocations []Allocation, cost float64)
	itemList := make([]string, 0, len(order))
	for item := range order {
		itemList = append(itemList, item)
	}

	backtrack = func(itemIndex int, allocations []Allocation, cost float64) {
		if cost >= minCost {
			return
		}

		if itemIndex == len(itemList) {
			for _, remaining := range order {
				if remaining > 0 {
					return
				}
			}
			successful = true
			if cost < minCost {
				minCost = cost
				optimalAllocations = make([]Allocation, len(allocations))
				copy(optimalAllocations, allocations)
			}
			return
		}

		item := itemList[itemIndex]
		quantityNeeded := order[item]

		for _, r := range restaurants {
			if price, ok := r.Menu[item]; ok {
				capacityLeft := r.MaxCapacity - r.CurrentCapacity
				if capacityLeft > 0 {

					qtyToAllocate := min(quantityNeeded, capacityLeft)
					r.CurrentCapacity += qtyToAllocate
					order[item] -= qtyToAllocate

					allocations = append(allocations, Allocation{
						RestaurantId: r.ID,
						Item:         item,
						Quantity:     qtyToAllocate,
						Cost:         float64(qtyToAllocate) * price,
					})

					backtrack(itemIndex+1, allocations, cost+float64(qtyToAllocate)*price)

					// backtracking
					r.CurrentCapacity -= qtyToAllocate
					order[item] += qtyToAllocate
					allocations = allocations[:len(allocations)-1]
				}
			}
		}
	}

	backtrack(0, []Allocation{}, 0)

	if !successful {
		return nil, 0, fmt.Errorf("unable to fulfill order")
	}

	return optimalAllocations, minCost, nil
}
