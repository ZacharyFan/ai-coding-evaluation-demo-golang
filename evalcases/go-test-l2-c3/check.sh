#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestRateRejectsEmptyShipments" "$ROOT/shipping/shipping_test.go" >/dev/null
cd "$ROOT"
go test ./shipping -run '^TestRateRejectsEmptyShipments$' -count=1
