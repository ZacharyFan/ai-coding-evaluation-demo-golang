package pricing

import "fmt"

func TaxCents(amountCents int, rateBasisPoints int) (int, error) {
	if amountCents < 0 {
		return 0, fmt.Errorf("amount cents must be non-negative")
	}
	if rateBasisPoints < 0 {
		return 0, fmt.Errorf("rate basis points must be non-negative")
	}
	return amountCents * rateBasisPoints / 10000, nil
}

func ApplyDiscount(amountCents int, percent int) (int, error) {
	if amountCents < 0 {
		return 0, fmt.Errorf("amount cents must be non-negative")
	}
	if percent < 0 || percent > 100 {
		return 0, fmt.Errorf("percent must be between 0 and 100")
	}
	return amountCents * (100 - percent) / 100, nil
}
