package main

import "fmt"

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.batteryDrain <= car.battery {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%s", car.battery, "%")
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	return car.battery/car.batteryDrain*car.speed >= trackDistance
}

func main() {
	car := NewCar(5, 2)
	fmt.Println(car)
	fmt.Println(car.DisplayDistance())
	fmt.Println(car.DisplayBattery())
	fmt.Println(car.CanFinish(100))
}
