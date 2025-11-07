package main

import "fmt"

func main() {
	fmt.Println(ToRomanNumeral(1996))
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", fmt.Errorf("input must be between 1 and 3999")
	}

	numberDict := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanDict := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string

	for i := 0; i < len(numberDict); i++ {
		for input >= numberDict[i] {
			input -= numberDict[i]
			result += romanDict[i]
		}
	}

	return result, nil
}
