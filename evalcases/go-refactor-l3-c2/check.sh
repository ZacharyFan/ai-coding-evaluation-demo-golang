#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func shouldRetryPaymentError" "$ROOT/payment_saga/payment.go" >/dev/null
grep -R "shouldRetryPaymentError(err)" "$ROOT/payment_saga/payment.go" >/dev/null
