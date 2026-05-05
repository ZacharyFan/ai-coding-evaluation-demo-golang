# Checkout Context

Checkout is a multi-domain monolith flow.

Domains:
- cart: validates order lines and computes subtotal.
- inventory: verifies every requested SKU has available stock.
- pricing: applies discounts and tax.
- shipping: computes shipping from item count and discounted subtotal.

Contract:
1. Validate all cart lines.
2. Validate all inventory lines before returning a quote.
3. Apply discounts before tax.
4. Shipping uses the discounted subtotal for free-shipping thresholds.
5. Total = discounted subtotal + tax + shipping.
