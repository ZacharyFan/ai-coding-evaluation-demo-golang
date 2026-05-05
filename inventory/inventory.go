package inventory

import "fmt"

type Snapshot map[string]int

func Available(snapshot Snapshot, sku string) int {
	if snapshot[sku] < 0 {
		return 0
	}
	return snapshot[sku]
}

func HasAvailable(snapshot Snapshot, sku string, quantity int) bool {
	return quantity > 0 && Available(snapshot, sku) >= quantity
}

func Restock(snapshot Snapshot, sku string, quantity int) error {
	if sku == "" {
		return fmt.Errorf("sku is required")
	}
	if quantity <= 0 {
		return fmt.Errorf("quantity must be positive")
	}
	snapshot[sku] += quantity
	return nil
}
