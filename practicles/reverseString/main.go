package main

import (
	"fmt"
)

func reverseString(str string) string {
	value := ""

	for _, v := range str {
		value = string(v) + value
	}

	return value
}

func main() {
	result := reverseString("Hello")
	fmt.Println(result)
	// fmt.Println(string(64))
}
