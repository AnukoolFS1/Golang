package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (err *MyError) Error() string{
	return fmt.Sprintf("at %v Error %s occured", err.when, err.what)
}

func run(x int) error {
	if x == 0 {
		return nil
	}
	return &MyError{
		time.Now(),
		"Class",
	}
}