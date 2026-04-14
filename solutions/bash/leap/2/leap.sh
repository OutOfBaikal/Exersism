#!/usr/bin/env bash

if [[ $# -ne 1 ]]; then
    echo "Usage: leap.sh <year>" >&2
    exit 1
fi

year=$1

if [[ "$year" =~ [^0-9] ]]; then
    echo "Usage: leap.sh <year>" >&2
    exit 1
fi 

if (( year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) )); then
    echo "true"
else
    echo "false"
fi
