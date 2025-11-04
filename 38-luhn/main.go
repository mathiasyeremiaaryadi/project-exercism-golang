package main

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	cleanedID := strings.ReplaceAll(id, " ", "")
	total := 0

	if len(cleanedID) <= 1 {
		return false
	}

	for index, num := range cleanedID {
		numInt, err := strconv.Atoi(string(num))
		if err != nil {
			return false
		}

		if len(cleanedID)%2 == 0 && index%2 == 0 {
			numInt *= 2
			if numInt > 9 {
				numInt -= 9
			}
		}

		if len(cleanedID)%2 != 0 && index%2 != 0 {
			numInt *= 2
			if numInt > 9 {
				numInt -= 9
			}
		}

		total += numInt
	}

	return total%10 == 0
}

func main() {
	println(Valid("4539 1488 0343 6467")) // true
	println(Valid("8273 1232 7352 0569")) // false
	println(Valid("1234567890123456"))    // false
	println(Valid("1234 5678 9012 3456")) // true
}
