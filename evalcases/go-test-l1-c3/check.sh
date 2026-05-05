#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestMaskEmailPreservesDomain" "$ROOT/account/account_test.go" >/dev/null
cd "$ROOT"
go test ./account -run '^TestMaskEmailPreservesDomain$' -count=1
