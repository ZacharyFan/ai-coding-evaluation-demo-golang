#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestReconcileReturnsNotificationFailure" "$ROOT/settlement_reconciliation/settlement_test.go" >/dev/null
cd "$ROOT"
go test ./settlement_reconciliation -run '^TestReconcileReturnsNotificationFailure$' -count=1
