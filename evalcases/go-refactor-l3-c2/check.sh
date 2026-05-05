#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func validateReservationRequest" "$ROOT/inventory" >/dev/null
grep -R "validateReservationRequest(l, sku, quantity)" "$ROOT/inventory" >/dev/null
