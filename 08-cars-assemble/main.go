package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int((float64(productionRate) * (successRate / 100)) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOf := uint(carsCount) / 10
	individualCars := uint(carsCount) % 10
	return (groupsOf * 95000) + (individualCars * 10000)
}

func main() {
	fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
	// => 1392.3

	fmt.Println(CalculateWorkingCarsPerMinute(1105, 90))
	// => 16

	fmt.Println(CalculateCost(37))
	// => 355000

	fmt.Println(CalculateCost(21))
	// => 200000

}
