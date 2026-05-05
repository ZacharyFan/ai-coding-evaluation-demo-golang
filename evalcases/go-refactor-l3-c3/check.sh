#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func matchPayoutsToLedger" "$ROOT/settlement_reconciliation/settlement.go" >/dev/null
grep -R "matchPayoutsToLedger(payouts, entries)" "$ROOT/settlement_reconciliation/settlement.go" >/dev/null
