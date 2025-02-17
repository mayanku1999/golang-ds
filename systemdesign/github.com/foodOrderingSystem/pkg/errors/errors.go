package errors

import (
	"errors"
	"fmt"
)

var (
	InValidMenuErr             = errors.New("invalid input: restaurantID or updatedMenu is empty")
	RestaurantNotFoundErr      = errors.New("restaurant not found")
	OrderNotFound              = errors.New("order not found")
	UndefinedSelectionStrategy = errors.New("undefined selection strategy")
	InvalidResNameErr          = errors.New("restaurant name cannot be empty")
	InvalidMenuPriceErr        = func(item string) error { return fmt.Errorf("invalid price for item: %s", item) }
	OrderCreationFailureErr    = func(err error) error { return fmt.Errorf("failed to create order: %w", err) }
	UpdatingCapErr             = func(restaurantId string, err error) error {
		return fmt.Errorf("failed to update capacity for restaurant %s: %w", restaurantId, err)
	}
)
