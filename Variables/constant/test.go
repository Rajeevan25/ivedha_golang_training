package main

import "fmt"

func main() {
	const pi float64 = 3.14159
	fmt.Println(pi)

	const (
		one   = "Monday"
		two   = "Tuesday"
		three = "Wednesday"
		four  = "Thursday"
		five  = "Friday"
		six   = "Saturday"
		seven = "Sunday"
	)
	fmt.Println("Days of the week:")
	fmt.Printf(" %v\n %v\n %v\n %v\n %v\n %v\n %v\n", one, two, three, four, five, six, seven)

	const hoursInDay float64 = 13
	const result float64 = hoursInDay * 60
	fmt.Println("The number of minutes in a day ", result)

	const taxRate float64 = 0.08
	fmt.Println("Tax for a purchase of $120 : ", taxRate*120)

	const version string = "17.1.2"
	fmt.Printf("Current version is %v\n", version)

}
