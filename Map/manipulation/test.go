package main

import "fmt"

func main() {
	ages := map[string]int{}
    ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28


	// scores := map[string]float64{
	// 	"maths" : 92.5,
	// 	"science": 83,
	// 	"history": 85,
	// 	"english": 75.3,
	// 	"commerce" : 67.8,
	// }

	fmt.Println(" value associated with the key Bob from the ages map",ages["Bob"])

	delete(ages,"Alice")
	fmt.Println("modified map : ",ages)
	ages["Charlie"] = 32
	fmt.Println("updated ages map : ",ages)
	

}
