package techpalace

import (
    "strings"
    "regexp"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    data := strings.Repeat("*",  numStarsPerLine)
    return data + "\n" + welcomeMsg + "\n" + data
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    cleanedMessage := regexp.MustCompile(`\*+`).ReplaceAllString(oldMsg, "")
	// Remove leading and trailing whitespaces from the remaining text
	trimmedMessage := strings.TrimSpace(cleanedMessage)
	return trimmedMessage
	panic("Please implement the CleanupMessage() function")
}
