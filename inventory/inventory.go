package inventory

import (
	"fmt"
	"sort"
)

type Item struct {
	SKU           string
	Stock         int
	Reserved      int
	ReorderPoint  int
	LeadTimeDays  int
	WarehouseCode string
}

type Ledger struct {
	stock    map[string]int
	reserved map[string]int
}

func Available(item Item) int {
	return item.Stock - item.Reserved
}

func NewLedger(stock map[string]int) *Ledger {
	copied := make(map[string]int, len(stock))
	for sku, quantity := range stock {
		copied[sku] = quantity
	}
	return &Ledger{
		stock:    copied,
		reserved: map[string]int{},
	}
}

func (l *Ledger) Reserve(sku string, quantity int) error {
	if sku == "" {
		return fmt.Errorf("sku is required")
	}
	if quantity <= 0 {
		return fmt.Errorf("quantity must be positive")
	}
	if l.stock[sku]-l.reserved[sku] < quantity {
		return fmt.Errorf("insufficient stock for %s", sku)
	}
	l.reserved[sku] += quantity
	return nil
}

func (l *Ledger) ReserveMany(requests map[string]int) error {
	keys := make([]string, 0, len(requests))
	for sku := range requests {
		keys = append(keys, sku)
	}
	sort.Strings(keys)
	for _, sku := range keys {
		if err := l.Reserve(sku, requests[sku]); err != nil {
			return err
		}
	}
	return nil
}

func (l *Ledger) Snapshot() map[string]int {
	result := make(map[string]int, len(l.stock))
	for sku, quantity := range l.stock {
		result[sku] = quantity - l.reserved[sku]
	}
	return result
}
