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

func TestApplyDiscountRejectsInvalidPercent(t *testing.T) {
	tests := []struct {
		name    string
		percent int
	}{
		{name: "negative", percent: -1},
		{name: "over one hundred", percent: 101},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := ApplyDiscount(1000, tt.percent); err == nil {
				t.Fatalf("ApplyDiscount(1000, %d) = %d, nil error; want error", tt.percent, got)
			}
		})
	}
}

func TestApplyDiscountRejectsNegativeTotal(t *testing.T) {
	if got, err := ApplyDiscount(-1, 10); err == nil {
		t.Fatalf("ApplyDiscount(-1, 10) = %d, nil error; want error", got)
	}
}
