package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var nums = [3]int{1, 2, 3}
	slice := nums[:]
	fmt.Println(slice)

	data := []interface{}{1, "Hello", true}
	fmt.Println(reflect.TypeOf(data[0]))
	fmt.Println(reflect.TypeOf(data[2]))

	for i , v := range data {
		if intValue, ok := v.(int); ok {
			data[i] =strconv.Itoa(intValue)
		}
		if boolValue, ok := v.(bool); ok {
			data[i] =strconv.FormatBool(boolValue)
		}		
	}	
	fmt.Println(reflect.TypeOf(data[0]))
	fmt.Println(reflect.TypeOf(data[2]))

	var str string = "Golang"
	byteSlice := []byte(str)
	fmt.Println(byteSlice)

}
