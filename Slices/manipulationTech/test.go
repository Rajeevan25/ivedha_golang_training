package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5}

	copiedNums := make([]int, len(numbers))
	copy(copiedNums, numbers)
	copiedNums = append(copiedNums, 11,12)
	fmt.Println("original slice : ",numbers)
	fmt.Println("copied  slice : ",copiedNums)

	names := []string{"Alice", "Bob", "Charlie"}
	names = names[1:]
	fmt.Println("modified slice : ",names)

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	mergedSlice := append(slice1,slice2...)
	fmt.Println("Merged slice : ", mergedSlice)
}
