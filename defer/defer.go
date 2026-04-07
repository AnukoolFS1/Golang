package main

import "fmt"

func practiceDefer() {
	deferingThing := func() {
		defer fmt.Println("done")
		fmt.Println("working...")
	}
	defer deferingThing()
	fmt.Println("Hello")
}

func main() {
	defer practiceDefer()
	fmt.Println("initiating...")
}
