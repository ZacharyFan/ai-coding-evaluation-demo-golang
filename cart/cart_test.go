package cart

import "testing"

func TestApplyDiscountValidPercent(t *testing.T) {
	got, err := ApplyDiscount(1000, 25)
	if err != nil {
		t.Fatalf("ApplyDiscount returned error: %v", err)
	}
	if got != 750 {
		t.Fatalf("ApplyDiscount(1000, 25) = %d, want 750", got)
	}
}

func TestApplyDiscountRejectsNegativeTotal(t *testing.T) {
	if got, err := ApplyDiscount(-1, 10); err == nil {
		t.Fatalf("ApplyDiscount(-1, 10) = %d, nil error; want error", got)
	}
}

func TestAddFee(t *testing.T) {
	got, err := AddFee(1000, 250)
	if err != nil {
		t.Fatalf("AddFee returned error: %v", err)
	}
	if got != 1250 {
		t.Fatalf("AddFee(1000, 250) = %d, want 1250", got)
	}
}

func TestCartTotal(t *testing.T) {
	got, err := CartTotal([]Item{
		{SKU: "A", UnitCents: 200, Quantity: 2},
		{SKU: "B", UnitCents: 350, Quantity: 1},
	})
	if err != nil {
		t.Fatalf("CartTotal returned error: %v", err)
	}
	if got != 750 {
		t.Fatalf("CartTotal = %d, want 750", got)
	}
}
