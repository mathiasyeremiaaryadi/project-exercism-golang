package main

import "fmt"

type Planet string

var listPlanet = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	earthSecond := 31557600.0

	if _, ok := listPlanet[planet]; !ok {
		return -1.0
	}

	return seconds / (listPlanet[planet] * earthSecond)
}

func main() {
	fmt.Println(Age(1000000000, "Earth"))
	fmt.Println(Age(1000000000, "Mars"))
}
 