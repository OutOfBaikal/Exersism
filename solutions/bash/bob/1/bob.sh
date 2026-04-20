#!/usr/bin/env bash

remark=$(echo "$1" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')

has_letters() {
    [[ "$1" =~ [[:alpha:]] ]]
}

is_shouting() {
    [[ "$1" == "${1^^}" ]] && has_letters "$1"
}

if [[ -z "$remark" ]]; then
    echo "Fine. Be that way!"
    exit 0
fi

if [[ "$remark" == *"?" ]]; then
    if is_shouting "$remark"; then
        echo "Calm down, I know what I'm doing!"
    else
        echo "Sure."
    fi
elif is_shouting "$remark"; then
    echo "Whoa, chill out!"
else
    echo "Whatever."
fi