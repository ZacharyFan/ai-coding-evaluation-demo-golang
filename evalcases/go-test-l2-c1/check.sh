#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestBuildQuoteCombinesInventoryTaxAndShipping" "$ROOT/checkout/checkout_test.go" >/dev/null
cd "$ROOT"
go test ./checkout -run '^TestBuildQuoteCombinesInventoryTaxAndShipping$' -count=1
