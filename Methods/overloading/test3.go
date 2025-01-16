package main

import "fmt"

type Vehicle interface {
	startEngine()
	stopEngine()
}

type Car struct {
	brand  string
	status string
}

func (c *Car) startEngine() {
	if c.status == "running" {
		fmt.Println("Engine is already started")
	} else {
		c.status = "running"
		fmt.Println("Starting the engine")
	}
}
func (c *Car) stopEngine() {
	if c.status == "stop" {
		fmt.Println("Engine is already stoped")
	} else {
		c.status = "stop"
		fmt.Println("Stoping the engine")
	}

}

func main() {
	car1 := &Car{brand: "Tesla", status: "stop"}
	car1.startEngine()
	car1.startEngine()
	car1.stopEngine()
	car1.stopEngine()

}
