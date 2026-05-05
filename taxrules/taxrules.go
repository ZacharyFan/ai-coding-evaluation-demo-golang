package taxrules

import (
	"fmt"
	"strings"
)

const MaxRateBasisPoints = 2500

type Rule struct {
	Region          string
	Category        string
	RateBasisPoints int
}

func ValidateRule(rule Rule) error {
	if strings.TrimSpace(rule.Region) == "" {
		return fmt.Errorf("region is required")
	}
	if strings.TrimSpace(rule.Category) == "" {
		return fmt.Errorf("category is required")
	}
	if rule.RateBasisPoints < 0 || rule.RateBasisPoints > MaxRateBasisPoints {
		return fmt.Errorf("rate basis points out of range")
	}
	return nil
}

func TaxCents(amountCents int, rateBasisPoints int) (int, error) {
	if amountCents < 0 {
		return 0, fmt.Errorf("amount cents must be non-negative")
	}
	if rateBasisPoints < 0 || rateBasisPoints > MaxRateBasisPoints {
		return 0, fmt.Errorf("rate basis points out of range")
	}
	return amountCents * rateBasisPoints / 10000, nil
}

func ApplyRule(amountCents int, rule Rule) (int, error) {
	if err := ValidateRule(rule); err != nil {
		return 0, err
	}
	tax, err := TaxCents(amountCents, rule.RateBasisPoints)
	if err != nil {
		return 0, err
	}
	return amountCents + tax, nil
}
