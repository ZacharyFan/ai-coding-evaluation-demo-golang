#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestValidateProductRejectsNegativePrice" "$ROOT/catalog/catalog_test.go" >/dev/null
cd "$ROOT"
go test ./catalog -run '^TestValidateProductRejectsNegativePrice$' -count=1
