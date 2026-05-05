package checkout

import (
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/cart"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/inventory"
	"testing"
)

func TestBuildQuoteCombinesInventoryTaxAndShipping(t *testing.T) {
	quote, err := BuildQuote(Request{Lines: []cart.Line{{SKU: "sku-1", UnitCents: 1000, Quantity: 2}}, Inventory: inventory.Snapshot{"sku-1": 2}, TaxBasisPoints: 1000, Expedited: true})
	if err != nil {
		t.Fatalf("BuildQuote returned error: %v", err)
	}
	if quote.SubtotalCents != 2000 || quote.TaxCents != 200 || quote.ShippingCents != 1450 || quote.TotalCents != 3650 {
		t.Fatalf("quote = %#v, want subtotal=2000 tax=200 shipping=1450 total=3650", quote)
	}
}
