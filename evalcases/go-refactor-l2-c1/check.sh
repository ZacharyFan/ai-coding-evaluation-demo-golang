#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func validateInventoryForLines" "$ROOT/checkout/checkout.go" >/dev/null
grep -R "validateInventoryForLines(request.Lines, request.Inventory)" "$ROOT/checkout/checkout.go" >/dev/null
