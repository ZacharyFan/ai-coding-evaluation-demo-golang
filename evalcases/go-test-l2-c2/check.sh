#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestCanTransitionRejectsBackwardTransitions" "$ROOT/orders/orders_test.go" >/dev/null
cd "$ROOT"
go test ./orders -run '^TestCanTransitionRejectsBackwardTransitions$' -count=1
