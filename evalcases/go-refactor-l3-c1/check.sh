#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func taxCentsFromBasisPoints" "$ROOT/pricing" >/dev/null
grep -R "taxCentsFromBasisPoints(subtotalCents, basisPoints)" "$ROOT/pricing" >/dev/null
grep -R "taxCentsFromBasisPoints(subtotalCents, taxBasisPoints)" "$ROOT/pricing" >/dev/null
