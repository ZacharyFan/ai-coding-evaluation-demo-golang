#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func compensateAuthorization" "$ROOT/fulfillment_saga/fulfillment.go" >/dev/null
grep -R "compensateAuthorization(ctx, c.Inventory, c.Payment" "$ROOT/fulfillment_saga/fulfillment.go" >/dev/null
