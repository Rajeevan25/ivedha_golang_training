package main

import (
	"fmt"
)
func getAverageAge(ages map[string]int) float64  {
	total := 0
	for _, v := range ages {
		total += v
	}
	average := float64(total) / float64(len(ages))
	return  average
}
type  Details map[string]string
func (d Details) displayInfo()  {
	for k, v := range d{
		fmt.Println("Key : ",k,", Value : ", v )
	}
}

func filterMap(m map[string]int, threshold int) map[string]int  {
	filterM := map[string]int{}
	for k, v := range m {
		if v > threshold{
			filterM[k] = v
		}
	}
	return filterM
}

func main() {
	ages := map[string]int{}
    ages["Alice"] = 25
	ages["Bob"] = 30
	ages["John"] = 30
	ages["Charlie"] = 28
    delete(ages,"Alice")
	ages["Charlie"] = 32
    fmt.Println(ages)
	fmt.Println(getAverageAge(ages))

    details := Details{
		"maths" : "001",
		"science": "002",
		"history": "003",
		"english": "004",
		"commerce" : "005",
	}
	details.displayInfo()

	data := map[string]int{
		"maths" : 92,
		"science": 83,
		"history": 85,
		"english": 75,
		"commerce" : 67,
	}
	fmt.Println(filterMap(data,80))


	
}
