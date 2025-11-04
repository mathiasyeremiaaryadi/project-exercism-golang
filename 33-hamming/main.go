package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	hammingDistance := 0

	if len(a) != len(b) {
		return hammingDistance, errors.New("a and b should have same length")
	}

	for index, _ := range a {
		if a[index] != b[index] {
			hammingDistance += 1
		}
	}

	return hammingDistance, nil
}

func main() {
	fmt.Println(Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"))
}
