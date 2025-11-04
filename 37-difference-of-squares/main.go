package main

import "fmt"

func SquareOfSum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}

	return total * total
}

func SumOfSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i * i
	}

	return total
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func main() {
	fmt.Println(Difference(10))
}  
