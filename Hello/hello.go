package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	text := greetings.Hello("Gladys")
	fmt.Println(text)
	message := WheelWonder("Anukool")

	fmt.Printf("This is the %v", message)
	fmt.Println(greetings.AnyNumer)
}
