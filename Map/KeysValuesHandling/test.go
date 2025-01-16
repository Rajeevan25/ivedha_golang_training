package main

import "fmt"
 
func countValues(m map[string]int, value int) int  {
	var count int = 0 
	for _, v := range m{
		if v == value{
			count += 1
		}
	}
	return count
}
func main() {
	ages := map[string]int{}
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28

	for key := range ages {
		fmt.Println(key)
	}

	scores := map[string]float64{
		"maths" : 92.5,
		"science": 83,
		"history": 85,
		"english": 75.3,
		"commerce" : 67.8,
	}

	for _,value := range scores {
		fmt.Println(value)
	}

	counts := map[string]int{
		"maths" : 23,
		"science": 23,
		"history": 45,
		"english": 23,
		"commerce" : 68,
	}
	fmt.Println("Number of 23 in map is ",countValues(counts,23))
}
