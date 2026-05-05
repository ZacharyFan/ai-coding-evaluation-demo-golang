#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestPlaceOrderPublishesOutboxEvent" "$ROOT/fulfillment_saga/fulfillment_test.go" >/dev/null
cd "$ROOT"
go test ./fulfillment_saga -run '^TestPlaceOrderPublishesOutboxEvent$' -count=1
