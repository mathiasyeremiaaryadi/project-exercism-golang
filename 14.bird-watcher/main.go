package main

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, v := range birdsPerDay {
		sum += v
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	startWeek := 0
	endWeek := 0

	if week == 1 {
		startWeek = 0
		endWeek = 6
	} else {
		startWeek = 7
		endWeek = 13
	}

	for i := startWeek; i <= endWeek; i++ {
		sum += birdsPerDay[i]
	}

	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 1; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i-1] += 1
	}

	return birdsPerDay
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(TotalBirdCount(birdsPerDay))
	// => 34

	birdsPerDay = []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(BirdsInWeek(birdsPerDay, 2))
	// => 12

	birdsPerDay = []int{2, 5, 0, 7, 4, 1}
	fmt.Println(FixBirdCountLog(birdsPerDay))
	// => [3 5 1 7 5 1]
}
