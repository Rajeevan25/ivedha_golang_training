package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func concatenate(strings ...string) string {
	paragraph := ""
	for _, str := range strings {
		paragraph += str
		paragraph += " "
	}
	return paragraph
}
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(sum(13, 22, 43))
	fmt.Println(sum(981, 672, 323))

	result := func(a int) int {
		return 2 * a
	}(4)
	fmt.Println(result)

	fmt.Println(concatenate("Hello", "World", "!"))
	fmt.Println(concatenate("I", "want", "to", "be"))
	fmt.Println(concatenate("interact", "with", "a", "computer"))
}
