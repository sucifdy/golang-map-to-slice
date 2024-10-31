package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	var result [][]string

	for key, value := range mapData {
		pair := []string{key, value}
		result = append(result, pair)
	}

	return result
}

func main() {
	exampleMap1 := map[string]string{"hello": "world", "John": "Doe", "age": "14"}
	fmt.Println(MapToSlice(exampleMap1))
}
