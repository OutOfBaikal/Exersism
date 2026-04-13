#!/usr/bin/env bash

allergens=("eggs" "peanuts" "shellfish" "strawberries" "tomatoes" "chocolate" "pollen" "cats")

SCORE="$1"
COMMAND="$2"
TARGET="$3"

get_list() {
    local res=()
    for i in "${!allergens[@]}"; do
        if (((SCORE >> i) & 1)); then
            res+=("${allergens[$i]}")
        fi
    done
    echo "${res[*]}"
}

case "$COMMAND" in
    "list")
        get_list
        ;;
    "allergic_to")
        is_allergic="false"
        for i in "${!allergens[@]}"; do
            if [[ "${allergens[$i]}" == "$TARGET" ]]; then
                if (((SCORE >> i) & 1)); then
                    is_allergic="true"
                fi
                break
            fi
        done
        echo "$is_allergic"
        ;;
esac


