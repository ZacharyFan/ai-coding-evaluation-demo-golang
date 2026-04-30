package cart

import "fmt"

// ApplyDiscount returns the total after applying a percentage discount.
func ApplyDiscount(totalCents int, percent int) (int, error) {
	if totalCents < 0 {
		return 0, fmt.Errorf("total cents must be non-negative")
	}
	if percent < 0 || percent > 100 {
		return 0, fmt.Errorf("discount percent must be between 0 and 100")
	}
	return totalCents * (100 - percent) / 100, nil
}
