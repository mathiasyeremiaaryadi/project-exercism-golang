package main

import "fmt"

func main() {
	keepResult := Keep([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	discardResult := Discard([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })

	fmt.Println("Keep even numbers:", keepResult)
	fmt.Println("Discard even numbers:", discardResult)
}

func Keep[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Discard[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if !predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
