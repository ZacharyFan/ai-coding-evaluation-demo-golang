#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func returnEligibilityReason" "$ROOT/returns/returns.go" >/dev/null
grep -R "returnEligibilityReason(request.Order)" "$ROOT/returns/returns.go" >/dev/null
