package fulfillment_saga

import (
	"context"
	"testing"
)

func TestPlaceOrderPublishesOutboxEvent(t *testing.T) {
	outbox := &testOutbox{}
	coordinator := Coordinator{Inventory: okInventory{}, Payment: okPayment{}, Carrier: okCarrier{}, Outbox: outbox}
	err := coordinator.PlaceOrder(context.Background(), Request{OrderID: "ord-1", IdempotencyKey: "key-1", Lines: []Line{{SKU: "sku-1", Quantity: 1}}, AmountCents: 1000})
	if err != nil {
		t.Fatalf("PlaceOrder returned error: %v", err)
	}
	if outbox.topic != "fulfillment.created" || outbox.key != "ord-1" {
		t.Fatalf("outbox = %s/%s, want fulfillment.created/ord-1", outbox.topic, outbox.key)
	}
}

type okInventory struct{}

func (okInventory) Reserve(context.Context, string, []Line) error { return nil }
func (okInventory) Release(context.Context, string) error         { return nil }

type okPayment struct{}

func (okPayment) Authorize(context.Context, string, int) error { return nil }
func (okPayment) Void(context.Context, string) error           { return nil }

type okCarrier struct{}

func (okCarrier) CreateShipment(context.Context, string, []Line) error { return nil }

type testOutbox struct{ topic, key string }

func (o *testOutbox) Publish(_ context.Context, topic string, key string) error {
	o.topic = topic
	o.key = key
	return nil
}
