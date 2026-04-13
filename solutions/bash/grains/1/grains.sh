#!/usr/bin/env bash

n=$1

calculate_square() {
    local num=$1
    if (( num<1 || num>64 )); then
        echo "Error: invalid input" >&2
        exit 1
    fi
    echo "2 ^ ($num - 1)" | bc
}

calculate_total() {
    echo "2 ^ 64 - 1" | bc
 }

case "$n" in 
    "total")
        calculate_total
        ;;
    *)
        if [[ "$1" =~ ^[0-9]+$ ]]; then
            calculate_square "$1"
        else
            echo "Error: invalid input" >&2
            exit 1
        fi
        ;;
esac