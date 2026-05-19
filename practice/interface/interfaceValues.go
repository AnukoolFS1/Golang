// package main

// import (
// 	"fmt"
// 	"math"
// )

// type I interface {
// 	M()
// }

// type Ti struct {
// 	S string
// }

// func (t *Ti) M() {
// 	fmt.Println(t.S)
// }

// type Fi float64

// func (f Fi) M() {
// 	fmt.Println(f)
// }

// func main() {
// 	var i I

// 	i = &Ti{"Hello"}
// 	describe(i)
// 	i.M()

// 	i = Fi(math.Pi)
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }


package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	describe(i)
	var t *T
	i = t
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
