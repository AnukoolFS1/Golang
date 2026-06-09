package main

import "fmt"

func main() {

	type Vertex struct {
		X string
		Y int
		Z bool
	}

	// var value = Vertex{"Black", 2, false}

	// var value2 = Vertex{X:"anukool", Y: 3}

	// // fmt.Printf("%+v\n", value)
	// // fmt.Printf("%v\n", value2)


	/// MAPS ///
	type VertexTwo struct {
		Lat, Long float64
	}

	var m map[string]VertexTwo

	m = make(map[string]VertexTwo)

	m["Bell Labs"] = VertexTwo{
		40.68433, -74.39967,
	}

	fmt.Println(m)
}

// need to comeback here
