#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "type ReceiptRenderer struct" "$ROOT/orders" >/dev/null
grep -R "ReceiptRenderer{}.Text" "$ROOT/orders" >/dev/null
