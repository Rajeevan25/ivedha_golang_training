package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7}
	subNumbers := numbers[1 : len(numbers)-1]
	fmt.Println(subNumbers)


	var names = []string{"Alice", "Bob", "Charlie"}
	subNames := names[:2]
	fmt.Println(subNames)

	subNumbers1 := numbers[len(numbers)-3:]
	fmt.Println(subNumbers1)
}
