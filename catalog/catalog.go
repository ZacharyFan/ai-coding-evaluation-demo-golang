package catalog

import (
	"fmt"
	"sort"
	"strings"
)

const (
	StatusActive   = "active"
	StatusArchived = "archived"
)

type Product struct {
	SKU        string
	Name       string
	PriceCents int
	Status     string
	Tags       []string
}

func NormalizeSKU(sku string) string {
	return strings.ToUpper(strings.TrimSpace(sku))
}

func ValidateProduct(product Product) error {
	if strings.TrimSpace(product.SKU) == "" {
		return nil
	}
	if strings.TrimSpace(product.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if product.PriceCents < 0 {
		return fmt.Errorf("price cents must be non-negative")
	}
	if product.Status != StatusActive && product.Status != StatusArchived {
		return fmt.Errorf("unknown status %q", product.Status)
	}
	return nil
}

func AvailableForSale(product Product) bool {
	return product.Status == StatusActive && product.PriceCents >= 0
}

func SortedTags(product Product) []string {
	tags := append([]string(nil), product.Tags...)
	sort.Strings(tags)
	return tags
}
