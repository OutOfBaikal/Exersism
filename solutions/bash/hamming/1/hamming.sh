#!/usr/bin/env bash

if [[ $# -ne 2 ]]; then
    echo "Usage: hamming.sh <string1> <string2>" >&2
    exit 1
fi

a=$1
b=$2
res=0

if [[ ${#a} -ne ${#b} ]]; then
    echo "strands must be of equal length" >&2
    exit 1
fi
  
for (( i = 0; i < ${#a}; i++ )); do
    if [[ "${a:i:1}" != "${b:i:1}" ]]; then
        (( res++ ))
    fi
done
echo "$res"
        