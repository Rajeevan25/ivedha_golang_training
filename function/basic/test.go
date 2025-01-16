package main

import (
	"fmt"
	"math"
)

func greet() {
	fmt.Println("Hello, World!")
}

func add(a int, b int) int {
	return a + b

}
func calculateCircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}
func main() {
	greet()
	sum := add(4, 4)
	fmt.Println(sum)
	area := calculateCircleArea(7)
	fmt.Println(area)
}
