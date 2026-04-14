package house

import(
    "strings"
)

var phrases = []string{
    "the house that Jack built.",
    "the malt\nthat lay in",
    "the rat\nthat ate",
    "the cat\nthat killed",
    "the dog\nthat worried",
    "the cow with the crumpled horn\nthat tossed",
    "the maiden all forlorn\nthat milked",
    "the man all tattered and torn\nthat kissed",
    "the priest all shaven and shorn\nthat married",
    "the rooster that crowed in the morn\nthat woke",
    "the farmer sowing his corn\nthat kept",
    "the horse and the hound and the horn\nthat belonged to",
}

func Verse(v int) string {
	var sb strings.Builder
    sb.WriteString("This is ")
    for i := v - 1; i >= 0; i-- {
        sb.WriteString(phrases[i])
        if i != 0 {
            sb.WriteString(" ")
        }
    }

    return sb.String()
}

func Song() string {
    verses := make([]string, 0, 12)
	for i := 1; i <= 12; i++ {
        verses = append(verses, Verse(i))
    }

    return strings.Join(verses, "\n\n")
}
