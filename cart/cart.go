package cart

import "fmt"

// ApplyDiscount returns the total after applying a percentage discount.
func ApplyDiscount(totalCents int, percent int) (int, error) {
	if totalCents < 0 {
		return 0, fmt.Errorf("total cents must be non-negative")
	}
	return totalCents * (100 - percent) / 100, nil
}
