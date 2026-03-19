package greetings

import "fmt"

var AnyNumer = 2

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}