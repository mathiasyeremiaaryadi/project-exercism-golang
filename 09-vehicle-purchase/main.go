package main

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var suggestedOption string
	if option1 > option2 {
		suggestedOption = option2
	} else {
		suggestedOption = option1
	}

	return suggestedOption + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= 10 {
		return originalPrice * 0.5
	}

	if age >= 3 && age < 10 {
		return originalPrice * 0.7
	}

	return originalPrice * 0.8
}

func main() {
	fmt.Println(NeedsLicense("car"))
	// => true

	fmt.Println(NeedsLicense("bike"))
	// => false

	fmt.Println(ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))
	// => "Toyota Corolla is clearly the better choice."

	fmt.Println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
	// => "Volkswagen Beetle is clearly the better choice."

	fmt.Println(CalculateResellPrice(1000, 1))
	// => 800

	fmt.Println(CalculateResellPrice(1000, 5))
	// => 700

	fmt.Println(CalculateResellPrice(1000, 15))
	// => 500

}
