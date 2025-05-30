package greetings

import "fmt"

// Returns a greeting for a named person
func Hello(name string) string {
	// Returns a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}