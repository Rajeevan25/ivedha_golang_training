package main

import "fmt"

func sumMatrix(matrix [][]int) int  {
	var sum int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum += matrix[i][j]
		}
	}
	return sum
}
func main() {
	first := []int{34, 26, 67}
	second := []int{68, 39, 47}
	third := []int{57, 19, 34}

	matrix := [][]int{first, second, third}
	fmt.Println(matrix)

	fmt.Println("sum of all elements in a 2D integer slice ",sumMatrix(matrix))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}


}
