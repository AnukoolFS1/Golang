package main

import "fmt"

func main() {

	type Vertex struct {
		X int
		Y int
	}

	var value = Vertex{1, 2}

	fmt.Printf("v%", value)

}

// need to comeback here
