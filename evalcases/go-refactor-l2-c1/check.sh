#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func copyIntMap" "$ROOT/inventory" >/dev/null
grep -R "copyIntMap(stock)" "$ROOT/inventory" >/dev/null
grep -R "copyIntMap(available)" "$ROOT/inventory" >/dev/null
