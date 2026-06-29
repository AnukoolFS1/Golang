package main

import "fmt"

func removeDuplicate[T comparable](s []T) []T {
	elements := make(map[T]struct{})
	elementsSet := []T{}

	for _, v := range s {
		if _, exists := elements[v]; !exists {
			elements[v] = struct{}{}
			elementsSet = append(elementsSet, v)
		}
	}

	return elementsSet
}

func main() {
	result := removeDuplicate([]int{1, 23, 2, 23, 1, 4, 4, 5})

	fmt.Println(result)
}