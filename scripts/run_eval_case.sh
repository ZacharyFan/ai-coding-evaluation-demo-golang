#!/usr/bin/env bash
set -euo pipefail

TASK_ID="${1:?usage: scripts/run_eval_case.sh <task-id>}"
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CASE_DIR="$ROOT/evalcases/$TASK_ID"
TEMP_ROOT="$ROOT/evaltmp"
TEMP_DIR="$TEMP_ROOT/$TASK_ID"

if [[ ! -d "$CASE_DIR" ]]; then
  echo "unknown eval case: $TASK_ID" >&2
  exit 2
fi

cleanup() {
  rm -rf "$TEMP_ROOT"
}
trap cleanup EXIT

cd "$ROOT"
packages=()
while IFS= read -r package; do
  packages+=("$package")
done < <(go list ./... | grep -v '/evaltmp/')
go test "${packages[@]}"

shopt -s nullglob
go_case_files=("$CASE_DIR"/*.go.txt)
if (( ${#go_case_files[@]} > 0 )); then
  mkdir -p "$TEMP_DIR"
  for source in "${go_case_files[@]}"; do
    target="$TEMP_DIR/$(basename "${source%.txt}")"
    cp "$source" "$target"
  done
  go test "./evaltmp/$TASK_ID" -count=1
fi

if [[ -x "$CASE_DIR/check.sh" ]]; then
  "$CASE_DIR/check.sh" "$ROOT"
fi
