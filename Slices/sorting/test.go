package main

import (
	"fmt"
	"sort"
)
func reverseSlice(nums []int) []int  {
	var revNums []int
	for i := len(nums)-1; i > -1 ; i-- {
		revNums = append(revNums, nums[i])
	}
	return revNums

}
func main() {
	var numbers = []int{1, 2, 3, 4,34,598,234,123, 5,6,7}
	for i := 0; i < len(numbers); i++ {
		var mini = numbers[i]
		var x = i
		for j := i; j < len(numbers); j++ {
			if numbers[j] < mini {
				mini = numbers[j]
				x = j
			}
		}
		var temp = numbers[i]
		numbers[i] = numbers[x]
		numbers[x] = temp
	}
	fmt.Println(numbers)
	var nums = []int{1, 2,45, 3, 94,123, 5}
    sort.Ints(nums)
	fmt.Println("Sort the numbers slice in ascending order",nums)

	indexFound := -1
	for i, v := range numbers {		
		if v == 7{
			indexFound = i
		}
	}
	if indexFound == -1{
		fmt.Println("Not found")
	} else{
		fmt.Println("The index of 7 is ",indexFound)
	}
    fmt.Print(reverseSlice(numbers))
}
