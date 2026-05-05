package checkout

import (
	"fmt"

	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/cart"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/inventory"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/pricing"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/shipping"
)

type Request struct {
	Lines          []cart.Line
	Inventory      inventory.Snapshot
	TaxBasisPoints int
	Expedited      bool
}

type Quote struct {
	SubtotalCents int
	TaxCents      int
	ShippingCents int
	TotalCents    int
}

func BuildQuote(request Request) (Quote, error) {
	if len(request.Lines) == 0 {
		return Quote{}, fmt.Errorf("at least one line is required")
	}
	if err := validateInventoryForLines(request.Lines, request.Inventory); err != nil {
		return Quote{}, err
	}
	subtotal, err := cart.Subtotal(request.Lines)
	if err != nil {
		return Quote{}, err
	}
	tax, err := pricing.TaxCents(subtotal, request.TaxBasisPoints)
	if err != nil {
		return Quote{}, err
	}
	shippingCents, err := shipping.Rate(cart.ItemCount(request.Lines), request.Expedited, subtotal)
	if err != nil {
		return Quote{}, err
	}
	return Quote{SubtotalCents: subtotal, TaxCents: tax, ShippingCents: shippingCents, TotalCents: subtotal + tax + shippingCents}, nil
}

func validateInventoryForLines(lines []cart.Line, snapshot inventory.Snapshot) error {
	for _, line := range lines {
		if !inventory.HasAvailable(snapshot, line.SKU, line.Quantity) {
			return fmt.Errorf("insufficient inventory for %s", line.SKU)
		}
	}
	return nil
}
