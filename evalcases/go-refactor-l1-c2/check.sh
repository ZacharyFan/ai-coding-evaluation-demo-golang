#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "BasisPointsDenominator" "$ROOT/pricing" >/dev/null
if grep -R "[^A-Za-z]10000[^0-9]" "$ROOT/pricing/pricing.go" | grep -v "BasisPointsDenominator" >/dev/null; then
  echo "pricing package still uses raw 10000 outside the named constant" >&2
  exit 1
fi
