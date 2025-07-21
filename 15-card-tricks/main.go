package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	tempSlice := slice

	for i, _ := range tempSlice {
		if i == index {
			return tempSlice[i]
		}
	}

	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	tempSlice := slice

	if index > (len(tempSlice)-1) || index < 0 {
		tempSlice = append(tempSlice, value)
		return tempSlice
	}

	tempSlice[index] = value
	return tempSlice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	tempSlice := slice
	tempValues := values

	if len(tempValues) == 0 {
		return tempSlice
	}

	tempValues = append(tempValues, tempSlice...)
	return tempValues
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	tempSlice := slice
	var newSlice []int

	if index > (len(tempSlice)-1) || index < 0 {
		return tempSlice
	}

	for i, v := range tempSlice {
		if i == index {
			continue
		}

		newSlice = append(newSlice, v)
	}

	return newSlice
}

func main() {
	cards := FavoriteCards()
	fmt.Println(cards)
	// Output: [2 6 9]

	card := GetItem([]int{1, 2, 4, 1}, 2) // card == 4
	fmt.Println(card)

	card = GetItem([]int{1, 2, 4, 1}, 10) // card == -1
	fmt.Println(card)

	index := 2
	newCard := 6
	cards = SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards)
	// Output: [1 2 6 1]

	index = -1
	newCard = 6
	cards = SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards)
	// Output: [1 2 4 1 6]

	slice := []int{3, 2, 6, 4, 8}
	cards = PrependItems(slice, 5, 1)
	fmt.Println(cards)
	// Output: [5 1 3 2 6 4 8]

	slice = []int{3, 2, 6, 4, 8}
	cards = PrependItems(slice)
	fmt.Println(cards)
	// Output: [3 2 6 4 8]

	cards = RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println(cards)
	// Output: [3 2 4 8]

	cards = RemoveItem([]int{3, 2, 6, 4, 8}, 11)
	fmt.Println(cards)
	// Output: [3 2 6 4 8]
}
