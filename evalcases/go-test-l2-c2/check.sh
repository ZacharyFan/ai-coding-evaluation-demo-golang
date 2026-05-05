#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestProcessReturnDoesNotRestockWhenLedgerFails" "$ROOT/returns/returns_test.go" >/dev/null
cd "$ROOT"
go test ./returns -run '^TestProcessReturnDoesNotRestockWhenLedgerFails$' -count=1
