package main

import "fmt"

func isEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
func calculateVolume(length, width, height float64) float64 {
	return length * width * height
}
func fullName(firstName, lastName string) string {
	return firstName + " " + lastName
}
func main() {
	result := isEven(0)
	fmt.Println(result)
	volume := calculateVolume(5.2, 3.7, 4.0)
	fmt.Println(volume)
	name := fullName("Rajee", "Yoga")
	fmt.Println(name)
}
