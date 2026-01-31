// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
    res := make([]rune, 0)
    is_letter := false
    for i := 0; i < len(s); i++ {
        if !is_letter {
            if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') || s[i] == '\'' {
                is_letter = true
                res = append(res, rune(strings.ToUpper(string(s[i]))[0]))
            }
        } else {
            if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') && s[i] != '\'' {
                is_letter = false
            }
        }
    }
	return string(res)
}
