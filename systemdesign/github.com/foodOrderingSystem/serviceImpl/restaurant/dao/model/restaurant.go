package model

import (
	"fmt"
	"math/rand"
)

type Restaurant struct {
	ID              string
	Name            string
	Menu            map[string]float64
	MaxCapacity     int
	CurrentCapacity int
}

func NewRestaurant() *Restaurant {
	restaurantID := fmt.Sprintf("REST-%d", rand.Intn(1000000))
	return &Restaurant{ID: restaurantID}
}

func (r *Restaurant) WithName(name string) *Restaurant {
	r.Name = name
	return r
}

func (r *Restaurant) WithMenu(menu map[string]float64) *Restaurant {
	r.Menu = menu
	return r
}

func (r *Restaurant) WithMaxCapacity(max int) *Restaurant {
	r.MaxCapacity = max
	return r
}

func (r *Restaurant) WithCurrentCapacity(cap int) *Restaurant {
	r.CurrentCapacity = cap
	return r
}
