#!/usr/bin/env bash

INPUT="$1"

if [[ -n "$INPUT" ]]; then
    echo "One for $INPUT, one for me."
else
    echo "One for you, one for me."
fi
