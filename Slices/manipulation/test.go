package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6,7)

	var names = []string{"Alice", "Bob", "Charlie"}
	
	extractedNumbers := numbers[2:5]
	fmt.Println(extractedNumbers)

    newNames := names[:2]
	fmt.Println(newNames)

	numbers[len(numbers)-1] = 10
	numbers[len(numbers)-2] = 9
	numbers[len(numbers)-3] = 8
	fmt.Println(numbers)

	
}
