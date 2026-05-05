#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "StatusPending" "$ROOT/orders" >/dev/null
grep -R "StatusDelivered" "$ROOT/orders" >/dev/null
grep -R "from == StatusPending" "$ROOT/orders" >/dev/null
