# Catalog Context

Products are a single-domain aggregate managed by the catalog package.

Required fields:
- SKU must be present after trimming whitespace.
- SKU comparison is case-insensitive and uses uppercase normalized values.
- Name must be present.
- PriceCents must be zero or positive.
- Status is either active or archived.

Examples:
- SKU ` sku-1 ` is normalized to `SKU-1`.
- Active products with matching tags are available for tag searches.
- Archived products are not available for tag searches.
