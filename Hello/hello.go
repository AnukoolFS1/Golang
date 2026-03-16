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

	fmt.Println(message)
}
