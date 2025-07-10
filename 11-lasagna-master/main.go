package main

import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(sliceOfLayers []string, averageTimeLayer int) int {
	layerSize := len(sliceOfLayers)
	if averageTimeLayer == 0 {
		return layerSize * 2
	}

	return layerSize * averageTimeLayer
}

// TODO: define the 'Quantities()' function
func Quantities(sliceOfLayers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, v := range sliceOfLayers {
		if v == "noodles" {
			noodles += 50
		}

		if v == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(firstIngredients []string, secondIngredients []string) {
	sizeFirstIngredients := len(firstIngredients)
	sizeSecondIngredients := len(secondIngredients)
	lastFirstIngredients := sizeFirstIngredients - 1
	lastSecondIngredients := sizeSecondIngredients - 1

	secondIngredients[lastSecondIngredients] = firstIngredients[lastFirstIngredients]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(qty []float64, portion int) []float64 {
	tempQty := make([]float64, len(qty))
	for i, v := range qty {
		tempQty[i] = v * float64(portion) / 2.0
	}

	return tempQty
}

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	PreparationTime(layers, 3)
	// => 18
	PreparationTime(layers, 0)
	// => 12

	Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})
	// => 100, 0.4

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}

	AddSecretIngredient(friendsList, myList)
	// myList => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"}

	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantities, 4)
	fmt.Println(scaledQuantities)
	// => []float64{ 2.4, 7.2, 21 }
}
