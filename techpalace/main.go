package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {

	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {

	border := strings.Repeat("*", numStarsPerLine)

	template := "%s\n%s\n%s"

	return fmt.Sprintf(template, border, welcomeMsg, border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

	msgWithoutStars := strings.ReplaceAll(oldMsg, "*", "")
	trimmedMsg := strings.Trim(msgWithoutStars, " \n")

	return trimmedMsg
}
