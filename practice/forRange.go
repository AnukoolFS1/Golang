package main

import "fmt"

func anotherFor() {
	// var slice = []string{"alice", "bob", "mario"}

	// for i, v := range slice {
	// 	fmt.Println(i, v)
	// }

	var mapper = map[string]int{"Alice": 90, "Bob": 10}

	for key, value := range mapper {
		fmt.Printf("%v, %d\n", key, value)
	}
}

/// break, break outer, continue

func forKeyword() {
	outerloop:
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		for j := 1; j < 9; j++ {
			fmt.Println(j)
			if j == 2 {
				break outerloop
			}
		}
	}
}
