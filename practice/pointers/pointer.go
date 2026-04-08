package main

import "fmt"

func PointerUnderstanding() {
	i := 42
	p := &i

	fmt.Println(p)
	fmt.Println(*p)

	*p = *p - 20

	fmt.Println(i)
	fmt.Println(*p)


	// their respective type
	fmt.Printf("Type of val: %T\n", i) // Output: int
	fmt.Printf("Type of ptr: %T\n", p) // Output: *int

	fmt.Println(1 << 63 -1)
}

func main() {
	// i, j := 42, 2701

	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i

	// p = &j         // point to j
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j
	PointerUnderstanding()
}
