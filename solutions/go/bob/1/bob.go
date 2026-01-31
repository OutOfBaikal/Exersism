// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
    "strings"
    "unicode"
)

func ContainsLetters(input string) bool {
    for _, r := range input {
        if unicode.IsLetter(r) {
            return true
        }
    }
    return false
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
    trimmedRemark := strings.TrimSpace(remark) // Trim spaces from the start and end

    if len(trimmedRemark) > 0 {
        if trimmedRemark[len(trimmedRemark)-1] == '?' {
            if trimmedRemark == strings.ToUpper(trimmedRemark) && ContainsLetters(trimmedRemark) {
                return "Calm down, I know what I'm doing!"
            }
            return "Sure."
        }
        if trimmedRemark == strings.ToUpper(trimmedRemark) && ContainsLetters(trimmedRemark) {
            return "Whoa, chill out!"
        }
    }
    if trimmedRemark == "" {
        return "Fine. Be that way!"
    }
    return "Whatever."
}
