#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "BasisPointsDenominator" "$ROOT/taxrules/taxrules.go" >/dev/null
