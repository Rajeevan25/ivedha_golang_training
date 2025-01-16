package main

import (
	"fmt"
)

func main() {
	data := [][2]string{{"a", "apple"}, {"b", "banana"}, {"c", "cherry"}}
	fruits := map[string]string{}
	for i := 0; i < len(data); i++ {
		fruits[data[i][0]] = data[i][1]
	}
	fmt.Println(fruits)

    info := map[string]interface{}{"name": "Alice", "age": 30, "isStudent": true}
    modifyInfo := map[string]string{}
	for key, value := range info {
		if strValue, ok := value.(string); ok {
			modifyInfo[key] = strValue
		}
	}
	fmt.Println(modifyInfo)

	days := map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednesday"} 
    var slice []int
	for key := range days {
		slice = append(slice, key)
	}
	fmt.Println(slice)
}
