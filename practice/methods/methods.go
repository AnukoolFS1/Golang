package main

import (
	"fmt"
)

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

type MyFloat int

func (f MyFloat) Square() MyFloat {
	return f * f
}

func main() {
	// v := Vertex{3,4}

	// fmt.Println(v.Abs())


	f := MyFloat(3)
	fmt.Println(f.Square())

}
