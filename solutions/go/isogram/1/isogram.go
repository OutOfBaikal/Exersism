package isogram

import "strings"

func IsIsogram(word string) bool {
    word = strings.ToLower(word)
	for i := 0; i < len(word) - 1; i++ {
        for j := i + 1; j < len(word); j++ {
            if word[i] == word[j] && int(word[i]) >= int('a') && int(word[i]) <= int('z') {
                return false
            }
        }
    }
    return true
}
