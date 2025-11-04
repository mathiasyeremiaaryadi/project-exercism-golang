package main

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	charDict := make(map[string]int)
	for _, char := range word {
		cleanedChar := strings.ToLower(string(char))
		if _, ok := charDict[cleanedChar]; ok && cleanedChar != "-" && cleanedChar != " " {
			return false
		}

		charDict[cleanedChar] += 1
	}

	return true
}

func main() {
	fmt.Println(IsIsogram("lumberjacks"))
	fmt.Println(IsIsogram("background"))
	fmt.Println(IsIsogram("downstream"))
	fmt.Println(IsIsogram("six-year-old"))
}
