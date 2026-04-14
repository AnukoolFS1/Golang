package main

import "fmt"

func main() {

	type Vertex struct {
		X string
		Y int
		Z bool
	}

	var value = Vertex{"Black", 2, false}

	var value2 = Vertex{"anukool", 3, true}

	fmt.Printf("%+v\n", value)
	fmt.Printf("%v\n", value2)

	fmt.Println(value.X)
	fmt.Println(value.Y)
	fmt.Println(value2.X)
	fmt.Println(value2.Y)

	fmt.Println(value.Z)
	fmt.Println(value2.Z)

}

// need to comeback here
