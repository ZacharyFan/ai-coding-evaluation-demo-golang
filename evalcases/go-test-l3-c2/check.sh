#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestSettleInvoiceRetriesTransientGatewayError" "$ROOT/payment_saga/payment_test.go" >/dev/null
cd "$ROOT"
go test ./payment_saga -run '^TestSettleInvoiceRetriesTransientGatewayError$' -count=1
