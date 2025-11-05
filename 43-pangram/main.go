package main

import (
	"fmt"
	"unicode"
)

func IsPangram(input string) bool {
	temp := make(map[rune]bool)

	for _, char := range input {
		char = unicode.ToLower(char)
		if char >= 'a' && char <= 'z' {
			temp[char] = true
		}
	}

	return len(temp) == 26
}

func main() {
	fmt.Println(IsPangram("The quick brown fox jumps over the lazy dog"))
}
