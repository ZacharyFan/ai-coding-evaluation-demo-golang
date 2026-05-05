package checkout

import (
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/cart"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/inventory"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/pricing"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/shipping"
)

func BuildDiscountedQuote(request Request, discountPercent int) (Quote, error) {
	for _, line := range request.Lines {
		if !inventory.HasAvailable(request.Inventory, line.SKU, line.Quantity) {
			return Quote{}, errInsufficientInventory(line.SKU)
		}
	}
	subtotal, err := cart.Subtotal(request.Lines)
	if err != nil {
		return Quote{}, err
	}
	discounted, err := pricing.ApplyDiscount(subtotal, discountPercent)
	if err != nil {
		return Quote{}, err
	}
	tax, err := pricing.TaxCents(discounted, request.TaxBasisPoints)
	if err != nil {
		return Quote{}, err
	}
	shippingCents, err := shipping.Rate(cart.ItemCount(request.Lines), request.Expedited, discounted)
	if err != nil {
		return Quote{}, err
	}
	return Quote{SubtotalCents: discounted, TaxCents: tax, ShippingCents: shippingCents, TotalCents: discounted + tax + shippingCents}, nil
}

func errInsufficientInventory(sku string) error {
	return &inventoryError{sku: sku}
}

type inventoryError struct{ sku string }

func (e *inventoryError) Error() string { return "insufficient inventory for " + e.sku }
