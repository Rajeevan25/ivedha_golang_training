package main

import "fmt"

func main() {
	var age int = 25
	fmt.Printf("Type of age is %T, and the value is %v", age, age)
	println("")

	var city string = "New York"
	var name string = "Alice"
	fmt.Println(name + " from " + city)

	var temperature float64 = 78.5
	var newTemp int = int(temperature)
	fmt.Println("oldOne is ", temperature, " , newOne is ", newTemp)

	var isRaining bool = true
	if isRaining {
		fmt.Println("Take an umbrella")
	}

	var a, b int = 23, 56
	fmt.Println("Before swap x is ", a, " y is ", b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("After swap x is ", a, " y is ", b)
}
