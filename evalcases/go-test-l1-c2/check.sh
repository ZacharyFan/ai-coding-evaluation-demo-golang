#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestValidateRuleRejectsRateAboveCap" "$ROOT/taxrules/taxrules_test.go" >/dev/null
cd "$ROOT"
go test ./taxrules -run '^TestValidateRuleRejectsRateAboveCap$' -count=1
