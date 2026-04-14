#!/usr/bin/env bash

num=$1
res=""

if (( num % 3 == 0 )); then
    res+="Pling"
fi
if (( num % 5 == 0 )); then
    res+="Plang"
fi
if (( num % 7 == 0 )); then
    res+="Plong"
fi
if [[ -z "$res" ]]; then
    echo "$num"
else
    echo "$res"
fi
