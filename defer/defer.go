package main

import "fmt"

func practiceDefer() {
	deferingThing := func() {
		defer fmt.Println("done")
		fmt.Println("working...")
	}
	defer deferingThing()
}

func main() {
	practiceDefer()
	fmt.Println("initiating...")
}
