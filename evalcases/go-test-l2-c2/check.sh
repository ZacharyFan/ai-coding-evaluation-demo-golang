#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestProcessReturnRecordsLedgerAndRestocksOnSuccess" "$ROOT/returns/returns_test.go" >/dev/null
cd "$ROOT"
go test ./returns -run '^TestProcessReturnRecordsLedgerAndRestocksOnSuccess$' -count=1
