#!/usr/bin/env bash

name=$1
number=$2

if ! [[ $number =~ ^-?[0-9]+$ ]]; then
    echo "invalid" >&2
    exit 1
fi

num=${number#-}
end=$(( num % 10 ))
tail=$(( num % 100 ))

if (( tail != 11 && tail != 12 && tail != 13 )); then
    case $end in
        1) echo "${name}, you are the ${number}st customer we serve today. Thank you!" ;;
        2) echo "${name}, you are the ${number}nd customer we serve today. Thank you!" ;;
        3) echo "${name}, you are the ${number}rd customer we serve today. Thank you!" ;;
        *) echo "${name}, you are the ${number}th customer we serve today. Thank you!" ;;
    esac
else
     echo "${name}, you are the ${number}th customer we serve today. Thank you!"
fi
