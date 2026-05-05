# Payment Saga Context

This document is intentionally partial.

Known:
- PaymentGateway captures and refunds funds.
- LedgerClient records capture and refund events.
- Payment gateway can return transient errors.

Missing:
- Idempotency after ledger failure is not specified.
- Retry policy is incomplete.
- Refund failure handling is not fully documented.
