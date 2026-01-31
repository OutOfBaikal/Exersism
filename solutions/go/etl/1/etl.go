package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int, 0)
    for k, v := range(in) {
        for i := 0; i < len(v); i++ {
            result[strings.ToLower(v[i])] = k
        }
    }
	return result
}
