#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestNetAfterReturnsRejectsOverReturn" "$ROOT/pricing/pricing_test.go" >/dev/null
cd "$ROOT"
go test ./pricing -run '^TestNetAfterReturnsRejectsOverReturn$' -count=1
