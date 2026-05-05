package catalog

import "testing"

func TestValidateProductAcceptsActiveProduct(t *testing.T) {
	err := ValidateProduct(Product{SKU: "sku-1", Name: "Notebook", PriceCents: 1200, Status: StatusActive})
	if err != nil {
		t.Fatalf("ValidateProduct returned error: %v", err)
	}
}

func TestNormalizeSKU(t *testing.T) {
	if got := NormalizeSKU(" sku-1 "); got != "SKU-1" {
		t.Fatalf("NormalizeSKU = %q, want SKU-1", got)
	}
}
