#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func validateProfileIdentity" "$ROOT/customerprofile/customerprofile.go" >/dev/null
grep -R "validateProfileIdentity(profile)" "$ROOT/customerprofile/customerprofile.go" >/dev/null
