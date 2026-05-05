#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestActivateCreatesInvoiceAndEntitlement" "$ROOT/subscriptions/subscriptions_test.go" >/dev/null
cd "$ROOT"
go test ./subscriptions -run '^TestActivateCreatesInvoiceAndEntitlement$' -count=1
