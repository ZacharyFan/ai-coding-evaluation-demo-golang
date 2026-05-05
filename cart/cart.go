package cart

import "fmt"

type Line struct {
	SKU       string
	UnitCents int
	Quantity  int
}

func Subtotal(lines []Line) (int, error) {
	total := 0
	for _, line := range lines {
		if line.SKU == "" {
			return 0, fmt.Errorf("sku is required")
		}
		if line.UnitCents < 0 {
			return 0, fmt.Errorf("unit cents must be non-negative")
		}
		if line.Quantity <= 0 {
			return 0, fmt.Errorf("quantity must be positive")
		}
		total += line.UnitCents * line.Quantity
	}
	return total, nil
}

func ItemCount(lines []Line) int {
	total := 0
	for _, line := range lines {
		total += line.Quantity
	}
	return total
}
