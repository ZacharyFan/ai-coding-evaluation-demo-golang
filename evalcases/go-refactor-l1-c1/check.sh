#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func validateNonNegativeCents" "$ROOT/cart" >/dev/null
grep -R 'validateNonNegativeCents("total cents"' "$ROOT/cart" >/dev/null
grep -R 'validateNonNegativeCents("unit cents"' "$ROOT/cart" >/dev/null
