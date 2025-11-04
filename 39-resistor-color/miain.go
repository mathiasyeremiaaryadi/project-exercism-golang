package main

import "fmt"

// Colors should return the list of all colors.
func Colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	var colorResistance int
	for index, c := range Colors() {
		if c == color {
			colorResistance = index
		}
	}
	return colorResistance
}

func main() {
	fmt.Println(Colors())
	fmt.Println(ColorCode("red"))
	fmt.Println(ColorCode("green"))
}
