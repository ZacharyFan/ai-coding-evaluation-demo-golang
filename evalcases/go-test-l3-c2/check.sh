#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestSubtotalAppliesLineDiscounts" "$ROOT/orders/orders_test.go" >/dev/null
cd "$ROOT"
go test ./orders -run '^TestSubtotalAppliesLineDiscounts$' -count=1
