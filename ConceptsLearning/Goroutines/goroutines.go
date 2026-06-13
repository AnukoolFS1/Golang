package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ { // range 5 or range ragne 1 << 1
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}