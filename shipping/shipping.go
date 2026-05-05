package shipping

import "fmt"

func Rate(itemCount int, expedited bool, discountedSubtotalCents int) (int, error) {
	if itemCount <= 0 {
		return 0, fmt.Errorf("item count must be positive")
	}
	if discountedSubtotalCents >= 10000 {
		return 0, nil
	}
	total := 500 + itemCount*100
	if expedited {
		total += 750
	}
	return total, nil
}
