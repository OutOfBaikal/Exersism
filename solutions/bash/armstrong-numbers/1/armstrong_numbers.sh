#!/usr/bin/env bash

sum=0
N=$1
len=${#N}

for (( i=0; i<len; i++)); do
    digit=${N:i:1}
    (( sum += digit ** len ))
done

if [[ $sum -eq $N ]]; then
    echo "true"
else
    echo "false"
fi

