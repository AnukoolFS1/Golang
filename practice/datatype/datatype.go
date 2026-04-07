package main

import "fmt"

func main() {
	var typeString string = "ABC"

	var typeInt int8 = 127 // any value between -2^7 to 2^7 - 1
	// int16 = combinational value between -2^15 to 2^15 - 1
	// int32 = combinational value between -2^31 to 2^31 - 1
	// int64 = combinational value between -2^63 to 2^63 - 1

	// uint8,16,32,64 = value from 0 to 2^(8 | 16 | 32 | 64) -1

	var Boolean bool = false

	fmt.Printf("value 1 %#v,\nvalue 2 %#v \n", typeString, typeInt)
	fmt.Printf("value 3 %#v\n", Boolean)

	fmt.Printf("Type 1: %T\n", typeString)
	fmt.Printf("Type 2: %T\n", typeInt)
	fmt.Printf("Type 3: %T\n", Boolean)
}
