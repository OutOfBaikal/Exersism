package bottlesong

import (
    "fmt"
    "strings"
)

// capitalizeFirstLetter capitalizes the first letter of a string.
func capitalizeFirstLetter(s string) string {
    if len(s) == 0 {
        return ""
    }
    return strings.ToUpper(string(s[0])) + s[1:]
}

func Recite(startBottles, takeDown int) []string {
	mapBottles := map[int]string{
        0:  "no",
        1:  "one",
        2:  "two",
        3:  "three",
        4:  "four",
        5:  "five",
        6:  "six",
        7:  "seven",
        8:  "eight",
        9:  "nine",
        10: "ten",
    }

    //var buffer strings.Builder
    endBottles := startBottles - takeDown + 1
    result := make([]string, 0)

    for i := startBottles; i >= endBottles; i-- {
        currentBottles := capitalizeFirstLetter(mapBottles[i])
        nextBottles := mapBottles[i-1]

        result = append(result, (fmt.Sprintf("%s green bottle%s hanging on the wall,", currentBottles, pluralize(i))))
        result = append(result, (fmt.Sprintf("%s green bottle%s hanging on the wall,", currentBottles, pluralize(i))))
        result = append(result, ("And if one green bottle should accidentally fall,"))
        result = append(result, (fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.", nextBottles, pluralize(i-1))))
        result = append(result, "")
    }
    return result[:len(result) - 1]
}

func pluralize(count int) string {
    if count == 1 {
        return ""
    }
    return "s"
}
