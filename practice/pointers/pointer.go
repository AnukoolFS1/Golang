package main

import (
	"fmt"
	"reflect"
)

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

	fmt.Println(1<<63 - 1)
}

func returnPointer() int {
	var value = 1
	fmt.Println(value)
	value += 1
	return value
}

func main() {
	// i := 42

	// fmt.Println(*returnPointer())
	returnedValue := returnPointer()

	fmt.Printf("%T\n", returnedValue)

	var x int = 42
	var y int8 = 100
	fmt.Println(reflect.TypeOf(x))        // int
	fmt.Println(reflect.TypeOf(y))        // int
	fmt.Println(reflect.TypeOf(x).Size()) // 8  (bytes)
	fmt.Println(reflect.TypeOf(y).Size()) // 8  (bytes)


	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i

	// p = &j         // point to j
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j
	// PointerUnderstanding()
}
