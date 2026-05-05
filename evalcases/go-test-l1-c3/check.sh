#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func TestCanSendMarketingRejectsDeletedProfiles" "$ROOT/customerprofile/customerprofile_test.go" >/dev/null
cd "$ROOT"
go test ./customerprofile -run '^TestCanSendMarketingRejectsDeletedProfiles$' -count=1
