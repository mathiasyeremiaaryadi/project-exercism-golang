package main

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	steps := 0

	if n <= 0 {
		return steps, errors.New("input must larger than 0")
	}

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			steps += 1
		} else {
			n = n*3 + 1
			steps += 1
		}
	}

	return steps, nil
}

func main() {
	fmt.Println(CollatzConjecture(12))
}
