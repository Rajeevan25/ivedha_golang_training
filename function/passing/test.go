package main

import "fmt"

func modifySlice(s []int) []int {
	for element := range s {
		s[element] *= 2
	}
	return s
}
func swap(x, y *int) {
	dum := *x
	*x = *y
	*y = dum
}
func main() {
	slice := []int{23, 45, 17}
	slice = modifySlice(slice)
	fmt.Println(slice)

	num1 := 12
	num2 := 23
	swap(&num1, &num2)
	fmt.Println("number 1 : ", num1)
	fmt.Println("number 2 : ", num2)

}
