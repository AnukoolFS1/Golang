package main

import (
	"fmt"
)

func findLargestElement(arr []int) int {
	if len(arr) == 0 {
		panic("empty slice")
	}
	largestElement := arr[0]
	for _, v := range arr{
		if largestElement < v{
			largestElement = v
		}
	}

	return largestElement
}

func main() {
	maxValue := findLargestElement([]int{5,1,2,3})
	fmt.Println(maxValue)
}