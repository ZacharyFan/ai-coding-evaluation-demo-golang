package cart

import "fmt"

type Item struct {
	SKU       string
	UnitCents int
	Quantity  int
}

func ApplyDiscount(totalCents int, percent int) (int, error) {
	if totalCents < 0 {
		return 0, fmt.Errorf("total cents must be non-negative")
	}
	return totalCents * (100 - percent) / 100, nil
}

func AddFee(totalCents int, feeCents int) (int, error) {
	if totalCents < 0 {
		return 0, fmt.Errorf("total cents must be non-negative")
	}
	if feeCents < 0 {
		return 0, fmt.Errorf("fee cents must be non-negative")
	}
	return totalCents + feeCents, nil
}

func LineTotal(unitCents int, quantity int) (int, error) {
	if unitCents < 0 {
		return 0, fmt.Errorf("unit cents must be non-negative")
	}
	if quantity <= 0 {
		return 0, fmt.Errorf("quantity must be positive")
	}
	return unitCents * quantity, nil
}

func CartTotal(items []Item) (int, error) {
	total := 0
	for _, item := range items {
		lineTotal, err := LineTotal(item.UnitCents, item.Quantity)
		if err != nil {
			return 0, err
		}
		total += lineTotal
	}
	return total, nil
}

func GiftWrapTotal(items []Item, feeEachCents int) (int, error) {
	if feeEachCents < 0 {
		return 0, fmt.Errorf("fee cents must be non-negative")
	}
	total := 0
	for _, item := range items {
		if item.UnitCents < 0 {
			return 0, fmt.Errorf("unit cents must be non-negative")
		}
		if item.Quantity <= 0 {
			return 0, fmt.Errorf("quantity must be positive")
		}
		total += item.Quantity * feeEachCents
	}
	return total, nil
}
