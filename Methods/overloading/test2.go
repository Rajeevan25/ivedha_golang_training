package main

import "fmt"

type Car struct {
	brand string
}

func (c Car) drive() {
	fmt.Println("The car is being driven")
}

type ElectricCar struct {
	Car             //embedding Car inside electicCar
	BatteryCapacity int
}

func main() {
	EleCar := ElectricCar{
		Car: Car{
			brand: "BMW",
		},
		BatteryCapacity: 85,
	}
	EleCar.drive()
}
