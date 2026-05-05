#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestReserveDoesNotMutateOnSingleSkuFailure" "$ROOT/inventory/inventory_test.go" >/dev/null
cd "$ROOT"
go test ./inventory -run '^TestReserveDoesNotMutateOnSingleSkuFailure$' -count=1
