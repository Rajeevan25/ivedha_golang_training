package main

import "fmt"


func containsString(strings []string, target string) bool  {
    result := -1
	for _, str := range strings {
		if str == target{
            result +=1
		}
	}
	if result == -1 {
		return false
	} else {
		return true
	}
}
type Numbers []int
func (n *Numbers) customAppend(item int, items ...int) []int  {
	*n = append(*n, item)
	*n = append(*n, items...)
	return *n
}
func uniqueSlice(nums []int) []int  {
	var uniqueNums []int
	for i := 0; i < len(nums); i++ {
		unique := false
		for j := i+1; j < len(nums); j++ {
			if nums[i] == nums[j]{
				unique = true
			}
		}
		if unique== false {
			uniqueNums = append(uniqueNums, nums[i])
		}	
	}
	return uniqueNums
	
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("if a given string exists in a slice of strings : ",containsString(names,"Bob"))
	fmt.Println("if a given string exists in a slice of strings : ",containsString(names,"Tharsi"))


	numbers := Numbers{1,2,3,4,5}
	numbers.customAppend(6,7)
	numbers.customAppend(8)
	numbers.customAppend(9,10,11,12,13)
	fmt.Println(numbers)

	var nums = []int{23,12,34,34,56,67,78,78,34}
	fmt.Println(uniqueSlice(nums))


}
