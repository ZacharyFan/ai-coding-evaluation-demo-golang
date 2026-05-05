# Returns Context

This document is intentionally partial.

Known happy path:
- Delivered orders may be returned.
- Refunds are calculated from the order total minus a restocking fee.
- Inventory receives returned units.
- Ledger records the refund.

Missing:
- Ledger failure ordering is not documented.
- Inventory rollback semantics are not documented.
- Preview behavior is not documented.
