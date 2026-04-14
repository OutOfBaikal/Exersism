#!/usr/bin/env bash

declare -A alphabet

input=$(echo "${1,,}")

for (( i=0; i<${#input}; i++ )); do
    char="${input:i:1}"
    if [[ "$char" =~ [a-z] ]]; then
        alphabet["$char"]=1
    fi
done

if [[ ${#alphabet[@]} -eq 26 ]]; then
    echo "true"
else
    echo "false"
fi
