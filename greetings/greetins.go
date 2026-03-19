package greetings

import (
	"errors"
	"fmt"
)

var AnyNumer = 2

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message, nil
}
