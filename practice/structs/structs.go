package main

import "fmt"

func main() {

	type Vertex struct {
		X string
		Y int
	}

	var value = Vertex{"1", 2}

	var value2 = Vertex{"anukool", 3}

	fmt.Printf("%+v\n", value)
	fmt.Printf("%v\n", value2)

}

// need to comeback here
