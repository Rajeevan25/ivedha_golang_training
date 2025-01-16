package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	numbers = append(numbers, 6,7)
	fmt.Println(numbers)

	var names = []string{"Alice", "Bob", "Charlie"}
	fmt.Printf("Length of slice is  %d, Capacity is  %d\n",len(names),cap(names))
}
