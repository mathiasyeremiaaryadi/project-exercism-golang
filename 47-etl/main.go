package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Transform(map[int][]string{1: {"A"}}))
}

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for key, charList := range in {
		for _, char := range charList {
			out[strings.ToLower(char)] = key
		}
	}

	return out
}
