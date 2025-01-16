package main

import "fmt"

func main() {
	ages := map[string]int{}
	fmt.Println(ages)

    ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28
	fmt.Println(ages)

	scores := map[string]float64{
		"maths" : 92.5,
		"science": 83,
		"history": 85,
		"english": 75.3,
		"commerce" : 67.8,
	}
	fmt.Println(len(scores))

}
