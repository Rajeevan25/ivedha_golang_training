package main

import "fmt"

func totalMarks(scores map[string]float64) float64  {
	var total float64 = 0 
	for _, s := range scores {
		total += s
	}
	return total
}
func main() {
	ages := map[string]int{}
    ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28
    delete(ages,"Alice")
	ages["Charlie"] = 32

	age , check := ages["Alice" ]
	if (check){
       fmt.Println("Age of Alice is ",age)
	}else {
		fmt.Println("Not Found")
	}
    fmt.Println(ages)
	for n, a := range ages {
		fmt.Println(n + "-", a)
	}

	scores := map[string]float64{
		"maths" : 92.5,
		"science": 83,
		"history": 85,
		"english": 75.3,
		"commerce" : 67.8,
	}
	fmt.Println("total marks from a given map of scores ",totalMarks(scores))

}
