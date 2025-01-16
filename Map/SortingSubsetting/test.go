package main

import (
	"fmt"
	"sort"
)


func main() {
	ages := map[string]int{}
	ages["Charlie"] = 28
	ages["Bob"] = 30
	ages["Alice"] = 25
	
	var keys []string
	for k := range ages {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Sorted Keys:", keys)

	subAges := map[string]int{}
	extract := []string{"Bob", "Charlie"}

	for _, k := range extract {
		if value,ok := ages[k]; ok {
			subAges[k] = value
		}
	}
	fmt.Println(subAges)
}
