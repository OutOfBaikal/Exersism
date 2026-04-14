#!/usr/bin/env bash

digits=$1
span=$2

if [[ $span -lt 0 ]]; then
    echo "span must not be negative" >&2
    exit 1 
fi
if [[ $span -gt ${#digits} ]]; then
    echo "span must not exceed string length" >&2
    exit 1
fi
if [[ "$digits" =~ [^0-9] ]]; then
    echo "input must only contain digits" >&2
    exit 1
fi
if [[ $span -eq 0 ]]; then
    echo 1
    exit 0
fi

res=0
for (( i=0; i<=${#digits}-span; i++ )); do
    current_product=1
    for (( j=0; j<$span; j++ )); do
        digit=${digits:i+j:1}
        (( current_product *= digit ))
    done
    if [[ $current_product -gt $res ]]; then
        res=$current_product
    fi
done
echo "$res"
        