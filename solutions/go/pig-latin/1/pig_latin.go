package piglatin

import (
    "strings"
)

func Sentence(sentence string) string {
	words := strings.Fields(sentence)
	for i, word := range words {
		words[i] = translateWord(word)
	}
	return strings.Join(words, " ")
}

func translateWord(word string) string {
	vowels := "aeiou"

	// Rule 1: Starts with vowel or "xr"/"yt"
	if strings.HasPrefix(word, "xr") || strings.HasPrefix(word, "yt") || strings.ContainsAny(string(word[0]), vowels) {
		return word + "ay"
	}

	// Rule 3: consonants + "qu"
	// Проверяем "qu" до обычных согласных, так как "qu" — это особый случай
	for i := 0; i < len(word)-1; i++ {
		if word[i:i+2] == "qu" {
			// Если перед qu нет гласных (или это начало слова)
			if !strings.ContainsAny(word[:i], vowels) {
				return word[i+2:] + word[:i+2] + "ay"
			}
			break // Если нашли гласную раньше qu, идем к Rule 2
		}
	}

	// Rule 4: consonants + "y"
	// Ищем "y" после одной или более согласных
	for i := 1; i < len(word); i++ {
		if word[i] == 'y' {
			if !strings.ContainsAny(word[:i], vowels) {
				return word[i:] + word[:i] + "ay"
			}
			break
		}
	}

	// Rule 2: Move consonants to the end
	for i, letter := range word {
		if strings.ContainsAny(string(letter), vowels) {
			return word[i:] + word[:i] + "ay"
		}
	}

	return word + "ay"
}
