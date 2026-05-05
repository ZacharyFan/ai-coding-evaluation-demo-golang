package refunds

import "fmt"

func Amount(totalCents int, restockingFeeCents int) (int, error) {
	if totalCents < 0 || restockingFeeCents < 0 {
		return 0, fmt.Errorf("amounts must be non-negative")
	}
	if restockingFeeCents > totalCents {
		return 0, fmt.Errorf("restocking fee cannot exceed total")
	}
	return totalCents - restockingFeeCents, nil
}
