package main

import "fmt"

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return prisonerIsAwake && !archerIsAwake
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (petDogIsPresent && !archerIsAwake) || (!petDogIsPresent && prisonerIsAwake && (!knightIsAwake && !archerIsAwake))
}

func main() {
	var knightIsAwake = false
	var archerIsAwake = true
	var prisonerIsAwake = false
	fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	// Output: true

	archerIsAwake = false
	prisonerIsAwake = true
	fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
	// Output: true

	knightIsAwake = false
	archerIsAwake = true
	prisonerIsAwake = false
	var petDogIsPresent = false
	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
	// Output: false
}
