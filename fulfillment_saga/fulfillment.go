package fulfillment_saga

import (
	"context"
	"fmt"
)

type Line struct {
	SKU      string
	Quantity int
}

type Request struct {
	OrderID        string
	IdempotencyKey string
	Lines          []Line
	AmountCents    int
}

type InventoryClient interface {
	Reserve(context.Context, string, []Line) error
	Release(context.Context, string) error
}

type PaymentClient interface {
	Authorize(context.Context, string, int) error
	Void(context.Context, string) error
}

type CarrierClient interface {
	CreateShipment(context.Context, string, []Line) error
}

type Outbox interface {
	Publish(context.Context, string, string) error
}

type Coordinator struct {
	Inventory InventoryClient
	Payment   PaymentClient
	Carrier   CarrierClient
	Outbox    Outbox
	processed map[string]struct{}
}

func (c *Coordinator) PlaceOrder(ctx context.Context, request Request) error {
	if request.IdempotencyKey == "" {
		return fmt.Errorf("idempotency key is required")
	}
	if err := c.Inventory.Reserve(ctx, request.IdempotencyKey, request.Lines); err != nil {
		return err
	}
	if err := c.Payment.Authorize(ctx, request.IdempotencyKey, request.AmountCents); err != nil {
		_ = c.Inventory.Release(ctx, request.IdempotencyKey)
		return err
	}
	if err := c.Carrier.CreateShipment(ctx, request.OrderID, request.Lines); err != nil {
		return err
	}
	if err := c.Outbox.Publish(ctx, "fulfillment.created", request.OrderID); err != nil {
		return err
	}
	return nil
}
