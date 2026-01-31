package anagram

import "strings"

func compareMaps(map1, map2 map[rune]int) bool {
    // Сравниваем длины карт
    if len(map1) != len(map2) {
        return false
    }

    // Сравниваем ключи и значения
    for key, value1 := range map1 {
        if value2, exists := map2[key]; !exists || value1 != value2 {
            return false
        }
    }

    return true
}

func Detect(subject string, candidates []string) []string {
	result := make([]string, 0)
    subject = strings.ToLower(subject)
    pattern := make(map[rune]int, 0)
    for _, v := range(subject) {
        pattern[v]++
    }
    for i := 0; i < len(candidates); i++ {
        word := strings.ToLower(candidates[i])
        if subject == word {
            continue
        }
        data := make(map[rune]int, 0)
        for _, v := range(word) {
        	data[v]++
    	}
        if compareMaps(pattern, data) {
            result = append(result, candidates[i])
        }
    }

    return result
}
