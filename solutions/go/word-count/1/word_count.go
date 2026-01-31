package wordcount

import (
    "strings"
    "regexp"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    result := make(Frequency)

    // Регулярное выражение для поиска слов, включая апострофы
    re := regexp.MustCompile(`\w+('\w+)?`)
    words := re.FindAllString(strings.ToLower(phrase), -1)

    for _, word := range words {
        result[word]++
    }

    return result
}
