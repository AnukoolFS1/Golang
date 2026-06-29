package main

import (
	"fmt"
	"time"
)

// 1. Define a function that takes a channel as an argument
func fetchUserData(userID int, ch chan string) {
	fmt.Printf("Fetching data for User %d...\n", userID)
	
	// Simulate a slow network request (2 seconds)
	time.Sleep(2 * time.Second) 
	
	result := fmt.Sprintf("User Profile Data for ID: %d", userID)
	
	// 2. Send the result into the channel
	ch <- result 
}

func main() {
	// 3. Create an unbuffered channel that handles strings
	dataChannel := make(chan string)

	fmt.Println("Main program starting...")

	// 4. Spin up the goroutine using the 'go' keyword
	go fetchUserData(42, dataChannel)

	fmt.Println("Main program is free to do other work while waiting!")
	fmt.Println("Processing checkout...")
	time.Sleep(500 * time.Millisecond) // Simulating light work

	fmt.Println("Main program is now ready for the user data. Waiting on channel...")
	
	// 5. RECEIVE from the channel. 
	// The main program will FREEZE right here until fetchUserData sends the data.
	// userData := <-dataChannel 

	fmt.Println("Received data successfully!")
	fmt.Println("Result:", <-dataChannel)
}