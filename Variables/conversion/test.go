package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Type

	var num1 int = 42
	fmt.Println(float64(num1))

	var num2 string = "42"
	intnum2, _ := strconv.Atoi(num2)
	fmt.Println("integer : ", intnum2)

	var weight float64 = 75.6
	var intWeight int = int(weight)
	fmt.Println("float value of weight is ", weight, " and integer value is ", intWeight)

	var status bool = true
	statusInt := 0
	if status {
		statusInt = 1
	}
	fmt.Println("Integer of the ", status, " is ", statusInt)

	//String

	var count int = 89
	fmt.Println("My marks is " + strconv.Itoa(count))

	var price string = "45"
	floatprice, _ := strconv.ParseFloat(price, 64)
	fmt.Println("Float price is ", floatprice)

	var distance float64 = 365.78
	dist := strconv.FormatFloat(distance, 'f', -1, 32)
	fmt.Println("Distance : " + dist)

	var isValid bool = true
	str := strconv.FormatBool(isValid)
	fmt.Println("The boolean value is " + str)

	var year int = 2023
	str2 := strconv.Itoa(year)
	fmt.Println("The current year is " + str2)

	var input string = "23"
	intInput, er := strconv.Atoi(input)
	if er != nil {
		fmt.Println("Invalid type error", er)
	}
	fmt.Printf("string change to integer : %v\n", intInput)

	var input1 string = "vijay"
	intInput1, er := strconv.Atoi(input1)
	if er != nil {
		fmt.Println("Invalid type error", er)
	} else {
		fmt.Printf("string change to integer : %v\n", intInput1)
	}

}
