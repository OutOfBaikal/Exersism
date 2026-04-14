#!/usr/bin/env bash

phrases=(
    "the house that Jack built."
    "the malt\nthat lay in"
    "the rat\nthat ate"
    "the cat\nthat killed"
    "the dog\nthat worried"
    "the cow with the crumpled horn\nthat tossed"
    "the maiden all forlorn\nthat milked"
    "the man all tattered and torn\nthat kissed"
    "the priest all shaven and shorn\nthat married"
    "the rooster that crowed in the morn\nthat woke"
    "the farmer sowing his corn\nthat kept"
    "the horse and the hound and the horn\nthat belonged to"
)
start=$1
end=$2

if [[ ! "$start" =~ ^[0-9]+$ || ! "$end" =~ ^[0-9]+$ ]] || \
   (( start < 1 || start > 12 || end < 1 || end > 12 || start > end )); then
    echo "invalid" >&2
    exit 1
fi

generate_verse() {
    local verse_idx=$1
    echo -n "This is "

    for (( i=verse_idx; i>=0; i-- )); do
        if (( i == 0 )); then
            echo -ne "${phrases[i]}"
        else
            echo -ne "${phrases[i]} "
        fi
    done
}

for (( j=start-1; j<end; j++ )); do
    generate_verse $j
    if (( j < end - 1 )); then
        echo -e "\n"
    else
        echo ""
    fi
done