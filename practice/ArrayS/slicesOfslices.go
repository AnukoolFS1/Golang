package main

import "strings"
import "fmt"

func SlicesOfSlice() {

	a := [][]string{
		[]string{"a", "b", "c"},
		{"1", "2", "3"}, // type is redundant
	}

	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(strings.Join(a[i], " "))
	// }
	for _, v:= range a{
		fmt.Println(strings.Join(v, "_"))
	}
}
