package error

import (
	"fmt"
	"errors"
)

func ErrorFunc(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}

	message := fmt.Sprintf("Hi, %v", name)
	return message, nil
}
