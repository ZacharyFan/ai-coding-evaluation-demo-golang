package pricing

import "testing"

func TestCalculateTaxExactCents(t *testing.T) {
	got, err := CalculateTax(10000, 825)
	if err != nil {
		t.Fatalf("CalculateTax returned error: %v", err)
	}
	if got != 825 {
		t.Fatalf("CalculateTax(10000, 825) = %d, want 825", got)
	}
}

func TestNetAfterReturns(t *testing.T) {
	got, err := NetAfterReturns(10000, 2500)
	if err != nil {
		t.Fatalf("NetAfterReturns returned error: %v", err)
	}
	if got != 7500 {
		t.Fatalf("NetAfterReturns = %d, want 7500", got)
	}
}

func TestInvoiceTotalWithoutCredit(t *testing.T) {
	got, err := InvoiceTotal(10000, 1000, 0)
	if err != nil {
		t.Fatalf("InvoiceTotal returned error: %v", err)
	}
	if got != 11000 {
		t.Fatalf("InvoiceTotal = %d, want 11000", got)
	}
}
