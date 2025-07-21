package main

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
		utfApp := fmt.Sprintf("%U", v)
		if utfApp == "U+2757" {
			return "recommendation"
		}

		if utfApp == "U+1F50D" {
			return "search"
		}

		if utfApp == "U+2600" {
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	result := ""
	for _, v := range log {
		if string(v) == string(oldRune) {
			result += string(newRune)
		} else {
			result += string(v)
		}
	}

	return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

func main() {
	fmt.Println(Application("❗ recommended search product 🔍"))
	// => recommendation

	log := "please replace '👎' with '👍'"
	fmt.Println(Replace(log, '👎', '👍'))
	// => please replace '👍' with '👍'"

	fmt.Println(WithinLimit("hello❗", 6))
	// => true
}
