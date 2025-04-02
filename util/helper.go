package util

import "fmt"

func CheckQuantity(stock, requested int) error {
	if stock-requested >= 0 {
		return nil
	}
	return fmt.Errorf("Insufficient stock. Available quantity: %d", stock)
}
