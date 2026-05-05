package orders

import "testing"

func TestSubtotal(t *testing.T) {
	got, err := Subtotal([]Line{
		{SKU: "A", UnitCents: 1000, Quantity: 2, DiscountPercent: 10},
		{SKU: "B", UnitCents: 500, Quantity: 1},
	})
	if err != nil {
		t.Fatalf("Subtotal returned error: %v", err)
	}
	if got != 2300 {
		t.Fatalf("Subtotal = %d, want 2300", got)
	}
}

func TestApplyCoupon(t *testing.T) {
	got, err := ApplyCoupon(10000, Coupon{Code: "SAVE10", PercentOff: 10})
	if err != nil {
		t.Fatalf("ApplyCoupon returned error: %v", err)
	}
	if got != 9000 {
		t.Fatalf("ApplyCoupon = %d, want 9000", got)
	}
}

func TestCanTransition(t *testing.T) {
	if !CanTransition("pending", "paid") {
		t.Fatal("pending should transition to paid")
	}
}

func TestPromotionDiscountCentsSinglePromotion(t *testing.T) {
	got, err := PromotionDiscountCents(Line{SKU: "A", UnitCents: 1000, Quantity: 2}, []Promotion{{SKU: "A", PercentOff: 10}})
	if err != nil {
		t.Fatalf("PromotionDiscountCents returned error: %v", err)
	}
	if got != 200 {
		t.Fatalf("PromotionDiscountCents = %d, want 200", got)
	}
}
