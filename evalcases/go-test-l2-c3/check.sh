#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestActivateRejectsHeldAccountWithoutSideEffects" "$ROOT/subscriptions/subscriptions_test.go" >/dev/null
cd "$ROOT"
go test ./subscriptions -run '^TestActivateRejectsHeldAccountWithoutSideEffects$' -count=1
