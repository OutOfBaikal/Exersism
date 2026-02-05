package foodchain

import (
	"strings"
)

var items = []struct {
	name    string
	phrase  string
}{
	{"fly", ""},
	{"spider", "It wriggled and jiggled and tickled inside her."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

func Verse(v int) string {
	if v < 1 || v > len(items) {
        return ""
    }

    var lines []string

    // 1. Вводная строка
    lines = append(lines, "I know an old lady who swallowed a "+items[v-1].name+".")

    // 2. Специальная фраза (если есть)
    if items[v-1].phrase != "" {
        lines = append(lines, items[v-1].phrase)
    }

    // 3. Цепочка проглатываний
    if v < len(items) {
        for i := v - 1; i > 0; i-- {
            curr := items[i].name
            prev := items[i-1].name

            var line string
            if prev == "spider" {
                // Добавляем описание паука как придаточную часть
                desc := items[i-1].phrase[3:len(items[i-1].phrase)-1]
                line = "She swallowed the " + curr + " to catch the " + prev + " that " + desc + "."
            } else {
                line = "She swallowed the " + curr + " to catch the " + prev + "."
            }
            lines = append(lines, line)
        }
        // Рефрен
        lines = append(lines, "I don't know why she swallowed the fly. Perhaps she'll die.")
    }

    return strings.Join(lines, "\n")
}

func Verses(start, end int) string {
	var verses []string
	for i := start; i <= end; i++ {
		verse := Verse(i)
		if verse != "" {
			verses = append(verses, verse)
		}
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, len(items))
}
