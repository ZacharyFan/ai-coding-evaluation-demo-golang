#!/usr/bin/env bash
set -euo pipefail
ROOT="${1:?}"
grep -R "func normalizeSKU" "$ROOT/catalog/catalog.go" >/dev/null
grep -R "normalizeSKU(product.SKU)" "$ROOT/catalog/catalog.go" >/dev/null
