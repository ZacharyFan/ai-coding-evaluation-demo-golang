package inventory

import "testing"

func TestReserveUpdatesAvailability(t *testing.T) {
	ledger := NewLedger(map[string]int{"sku-1": 5})
	if err := ledger.Reserve("sku-1", 2); err != nil {
		t.Fatalf("Reserve returned error: %v", err)
	}
	if got := ledger.Snapshot()["sku-1"]; got != 3 {
		t.Fatalf("available stock = %d, want 3", got)
	}
}

func TestReserveRejectsInsufficientStock(t *testing.T) {
	ledger := NewLedger(map[string]int{"sku-1": 1})
	if err := ledger.Reserve("sku-1", 2); err == nil {
		t.Fatal("Reserve returned nil error; want insufficient stock error")
	}
}

func TestAvailablePositiveStock(t *testing.T) {
	got := Available(Item{SKU: "sku-1", Stock: 10, Reserved: 3})
	if got != 7 {
		t.Fatalf("Available = %d, want 7", got)
	}
}
