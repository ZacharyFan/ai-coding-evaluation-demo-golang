#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func remoteAreaSurchargeCents" "$ROOT/shipping" >/dev/null
grep -R "remoteAreaSurchargeCents(shipment)" "$ROOT/shipping" >/dev/null
