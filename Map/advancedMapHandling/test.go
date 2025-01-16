package main

import "fmt"

func mergeMaps(map1, map2 map[string]int) map[string]int {
	mergeMap := map[string]int{}
	for key, value := range map1 {
		mergeMap[key] = value
	}
	for key, value := range map2 {
		if _, ok := mergeMap[key]; !ok {
			mergeMap[key] = value
		}
	}
	return mergeMap
}

func incrementMapValues(m map[string]int, increment int) {
	for k, v := range m {
		m[k] = v +  increment
	}
}

func extractUniqueValues(m map[int]string) []string  {
	uniqueValues := map[string]int{}
	for k, v := range m {
		uniqueValues[v] = k
	}
	var result []string
	for uniqueValue := range uniqueValues {
		result = append(result, uniqueValue)
	}
	return result

}

func main() {
	data1 := map[string]int{
		"maths":   92,
		"science": 83,
		"history":  85,
	}
	data2 := map[string]int{
		"history":  85,
		"english":  75,
		"commerce": 67,
	}
	mergeMap := mergeMaps(data1,data2)
	fmt.Println(mergeMap)

	incrementMapValues(mergeMap,5)
	fmt.Println(mergeMap)


    inputMap := map[int]string{
		1: "apple",
		2: "orange",
		3: "banana",
		4: "apple",
		5: "grape",
		6: "orange",
	}
	uniqueValues := extractUniqueValues(inputMap)
	fmt.Println("Unique Values:", uniqueValues)


}
