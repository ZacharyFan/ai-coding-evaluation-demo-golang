#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func activationDecision" "$ROOT/subscriptions/subscriptions.go" >/dev/null
grep -R "activationDecision(request.Account, request.Plan)" "$ROOT/subscriptions/subscriptions.go" >/dev/null
