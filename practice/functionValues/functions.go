package main

import (
	"fmt"
	// "math"
)

func compute(fn func(float64, float64) float64) float64 {

	return fn(3, 4)
}

func testing() (int, int) {
	return 1, 2
}

func closureFunction() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

var x, y int = testing()

func main() {
	inc := closureFunction() // now somewhere in memory the sum variable has shifted to heap

	value1 := inc(5)
	value2 := inc(2)

	fmt.Println(value1, value2)
	// hypot := func(x, y float64) float64 {
	// 	return math.Sqrt(x*x + y*y)
	// }
	// fmt.Println(hypot(5, 12))

	// fmt.Println(compute(hypot))
	// fmt.Println(compute(math.Pow))
}
