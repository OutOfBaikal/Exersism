#!/usr/bin/env bash
dna=$1
res=()
for (( i=0; i<${#dna}; i++ )); do
    char="${dna:$i:1}"
    if [[ "$char" == "G" ]]; then
        res+="C"
    elif [[ "$char" == "C" ]]; then
        res+="G"
    elif [[ "$char" == "T" ]]; then
        res+="A"
    elif [[ "$char" == "A" ]]; then
        res+="U"
    else
        echo "Invalid nucleotide detected." >&2
        exit 1
    fi
done
 
echo "$res"
