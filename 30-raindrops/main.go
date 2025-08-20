package main

import "fmt"

func Convert(number int) string {
	outputString := ""

	if number%3 == 0 {
		outputString += "Pling"
	}

	if number%5 == 0 {
		outputString += "Plang"
	}

	if number%7 == 0 {
		outputString += "Plong"
	}

	if outputString == "" {
		return fmt.Sprintf("%d", number)
	}

	return outputString
}

func main() {
	fmt.Println(Convert(28))
	fmt.Println(Convert(30))
	fmt.Println(Convert(34))
}
