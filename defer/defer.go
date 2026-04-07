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

func deferLIFO() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	defer practiceDefer()
	fmt.Println("initiating... π")
	deferLIFO()
}
