#!/bin/bash
set -ueo pipefail

status=$(git status --porcelain)
if [ -n "$status" ]; then
    echo "detected git diff after running generate; please re-run tasks"
    echo "$status"
    exit 1
fi