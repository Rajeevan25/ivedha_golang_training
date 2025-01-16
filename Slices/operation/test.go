package main

import "fmt"

func sumSlice(nums []int) int  {
    var sum int = 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func findMax(nums []int) int  {
    var max int = 0
	for _, v := range nums {
		if v > max{
			max = v
		}
	}	
	return max
}

func filterEven(nums []int) []int {
	var evenNums  []int
    for _, v := range nums {
		if v % 2 == 0 {
			evenNums = append(evenNums,v)
		}
	}
	return evenNums
}
func main() {
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of numbers slice is ", sumSlice(numbers))
	fmt.Println("Maximum of numbers slice is ", findMax(numbers))
    fmt.Println("Even numbers of numbers slice is ", filterEven(numbers))

	
}
