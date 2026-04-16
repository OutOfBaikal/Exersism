#!/usr/bin/env bash

declare -A colors=([black]=0 [brown]=1 [red]=2 [orange]=3 [yellow]=4 [green]=5 [blue]=6 [violet]=7 [grey]=8 [white]=9)

color1=${1:-}
color2=${2:-}

if [[ -z $color1 || -z $color2 ]]; then
    echo "invalid color" >&2
    exit 1
fi
if [[ -v colors["$color1"] && -v colors["$color2"] ]]; then
    res=$(( colors[$color1] * 10 + colors[$color2] ))
    echo "$res"
else
    echo "invalid color" >&2
    exit 1
fi

