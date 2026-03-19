package main

import (// just why does it not use comma to seperate imports 
	"example.com/greetings"
	"fmt"
	"log"
)



func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}

// func main() {
// 	// Get a greeting message and print it.
// 	text := greetings.Hello("Gladys")
// 	fmt.Println(text)
// 	message := WheelWonder("Anukool")

// 	fmt.Printf("This is the %v", message)
// 	fmt.Println(greetings.AnyNumer)
// }

/// Printf for dynamic string of changes
/// Sprintf for Printf but with returned value
/// Println just for console check
