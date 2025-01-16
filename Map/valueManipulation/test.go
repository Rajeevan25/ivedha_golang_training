package main

import (
	"fmt"

)

func main() {
	ages := map[string]int{}
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["John"] = 30
	ages["Charlie"] = 28
	delete(ages, "Alice")
	ages["Charlie"] = 32

	ages["Bob"] +=1
	fmt.Println(ages)
    
	counters := map[string]int{}
	counters["apple"] += 1
	fmt.Println(counters)


    _,ok := counters["orange"]  
	if (!ok){
		counters["orange"] = 1
	}
	fmt.Println(counters)


}
