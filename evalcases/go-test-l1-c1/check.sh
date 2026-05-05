#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestLineTotalRejectsInvalidQuantity" "$ROOT/cart/cart_test.go" >/dev/null
cd "$ROOT"
go test ./cart -run '^TestLineTotalRejectsInvalidQuantity$' -count=1
