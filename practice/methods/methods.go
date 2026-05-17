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

func (f MyFloat) Square(sum int) MyFloat {
	return (f * f) + MyFloat(sum)
}

type Vertex struct {
	x, y int
}

func (v Vertex) SeeResults() Vertex {
	v.x = 10
	v.y = 11

	return v
}

func (v *Vertex) ChangeResults() Vertex {
	v.x = 20
	v.y = 21

	return *v
}

func main() {
	// v := Vertex{3,4}
	// fmt.Println(v.Abs())

	// f := MyFloat(3)
	// fmt.Println(f.Square(5))

	var vertex = Vertex{1, 2}

	fmt.Println(vertex)

	var exp1 = vertex.SeeResults()
	fmt.Println(exp1)
	fmt.Println(vertex)

	var exp2 = vertex.ChangeResults()
	fmt.Println(exp2)
	fmt.Println(vertex)

}
