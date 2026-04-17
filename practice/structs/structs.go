package main

import "fmt"

func main() {

	type Vertex struct {
		X string
		Y int
		Z bool
	}

	var value = Vertex{"Black", 2, false}

	var value2 = Vertex{X:"anukool", Y: 3}

	fmt.Printf("%+v\n", value)
	fmt.Printf("%v\n", value2)

	// fmt.Println(value.X)
	// fmt.Println(value.Y)
	// fmt.Println(value2.X)
	// fmt.Println(value2.Y)

	// fmt.Println(value.Z)
	// fmt.Println(value2.Z)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

}

// need to comeback here
