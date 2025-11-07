package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Total grains: ", Total())
}

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("square number must be between 1 and 64")
	}

	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var result uint64
	for i := 1; i <= 64; i++ {
		grain, _ := Square(i)
		result += grain
	}

	return result
}
