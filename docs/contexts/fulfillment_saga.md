# Fulfillment Saga Context

Fulfillment crosses service boundaries through interfaces:
- InventoryClient reserves and releases stock.
- PaymentClient authorizes and voids payment.
- CarrierClient creates shipments.
- Outbox publishes durable workflow events.

Required order:
1. Reserve inventory.
2. Authorize payment.
3. Create shipment.
4. Publish `fulfillment.created` to the outbox.

Compensation:
- If payment authorization fails, release inventory.
- If carrier creation fails, void payment and release inventory.

Idempotency:
- IdempotencyKey is required.
- A completed key must not repeat reserve, authorize, or ship calls.
