#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func remainingCreditCents" "$ROOT/account" >/dev/null
grep -R "remainingCreditCents(account)" "$ROOT/account" >/dev/null
