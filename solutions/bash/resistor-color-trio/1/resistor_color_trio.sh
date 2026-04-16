#!/usr/bin/env bash

declare -A colors=([black]=0 [brown]=1 [red]=2 [orange]=3 [yellow]=4 [green]=5 [blue]=6 [violet]=7 [grey]=8 [white]=9)

color1=${1:-}
color2=${2:-}
color3=${3:-}

if [[ -z $color1 || -z $color2 || -z $color3 ]]; then
    echo "invalid color" >&2
    exit 1
fi
if [[ -v colors["$color1"] && -v colors["$color2"] && -v colors["$color3"] ]]; then
    res=$(( (colors[$color1] * 10 + colors[$color2]) * (10 ** colors[$color3]) ))
    buf=$res
    a=0
    while [[ $buf -ge 1000 && $a -lt 3 && $(( buf / 1000 * 1000 )) -eq $buf ]]; do
        (( buf /= 1000 ))
        (( a++ ))
    done
    ohms="ohms"
    if [[ $a -eq 1 ]]; then
        ohms="kiloohms"
    elif [[ $a -eq 2 ]]; then
        ohms="megaohms"
    elif [[ $a -eq 3 ]]; then
        ohms="gigaohms"
    fi
    echo "$buf $ohms"
else
    echo "invalid color" >&2
    exit 1
fi