#!/usr/bin/env bash

MODE=$1
N=$2

square_of_sum() {
    local res=0
    for (( i=1; i<=N; i++ )); do
        (( res += i))
    done
    echo $(( res * res ))
}

sum_of_squares() {
    local res=0
    for (( i=1; i<=N; i++ )); do
        (( res += i * i))
    done
    echo $res
}

case "$MODE" in
    "square_of_sum")
        square_of_sum
        ;;
    "sum_of_squares")
        sum_of_squares
        ;;
    "difference")
        val1=$(square_of_sum)
        val2=$(sum_of_squares)
        echo $(( val1 - val2 ))
        ;;
esac