package main

import "fmt"

func mergeSlices(slices ...[]int) []int {
	var mergeSlice []int
	for _, s := range slices {
		mergeSlice = append(mergeSlice, s...)
	}
	return mergeSlice

}
func splitSlice(nums []int, chunkSize int)  [][]int  {
	var newSlice [][]int
	for i := 0; i < len(nums); i += chunkSize {
		j := i + chunkSize
		if j > len(nums) {
			j = len(nums)
		}
		newSlice = append(newSlice, nums[i:j])
	}
	return newSlice
}

func padSlice(nums []int, length int) []int {
	
	if length <= len(nums) {
		return nums
	}
	newSlice := make([]int, len(nums), length)
	copy(newSlice, nums)

	for i := len(nums); i < length; i++ {
		newSlice = append(newSlice, 0)
	}
	return newSlice

}
func main() {
	slice1 := []int{41, 56, 98}
	slice2 := []int{54, 75, 80}
	slice3 := []int{47, 29, 38}

	fmt.Println(mergeSlices(slice1,slice2))
	fmt.Println(mergeSlices(slice1,slice3))
	fmt.Println(mergeSlices(slice2,slice3))
	mergeSlice := mergeSlices(slice1,slice2,slice3)
	fmt.Println(mergeSlice)

    fmt.Println(splitSlice(mergeSlice,2))

    fmt.Println(padSlice(slice1,4))

}
