package main

import "fmt"

// CurrentCondition of the city name.
var CurrentCondition string

// CurrentLocation of the weather.
var CurrentLocation string

// Forecast the currect location and its weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

func main() {
	condition := "sunny"
	city := "Jakarta"

	fmt.Println(Forecast(city, condition))
}
