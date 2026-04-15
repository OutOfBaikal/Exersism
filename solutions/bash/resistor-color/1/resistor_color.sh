#!/usr/bin/env bash

declare -A colors=(
  [black]=0 [brown]=1 [red]=2 [orange]=3 [yellow]=4
  [green]=5 [blue]=6 [violet]=7 [grey]=8 [white]=9
)

mode=${1:-}

case "$mode" in
    colors)
        order=("black" "brown" "red" "orange" "yellow" "green" "blue" "violet" "grey" "white")
        printf '%s\n' "${order[@]}"
        ;;
    code)
        color=${2:-}
        if [[ -z $color ]]; then
            echo "usage: $0 code <color>" >&2
            exit 1
        fi
        if [[ -v colors["$color"] ]]; then
            printf '%s\n' "${colors[$color]}"
        else
            echo "unknown color: $colors" >&2
            exit 1
        fi
        ;;
    *)
        echo "usage: $0 {colors|code <color>}" >&2
        exit 1
        ;;
esac
