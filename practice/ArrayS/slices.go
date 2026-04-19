package main

import "fmt"

// defining array in GO 
// var variableName = [length]type{el, el1, el2...} /// type of this is [l]T

func testingSlice() {
	slicedata := []int{1, 2, 3, 4, 5, 6}
	slicePointer := &slicedata

	fmt.Printf("slicePointer %p\n", slicePointer)

	fmt.Printf("slicePointer %p\n", slicePointer)
}

func main() {

	// var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	// var slicedArr1 = arr[4:]
	// var slicedArr2 = arr[:4]

	// fmt.Printf("%v\n", slicedArr1)
	// fmt.Printf("%v\n", slicedArr2)
	// testingSlice()
	SlicesOfSlice()
	fmt.Println(nil)


}
// func main() {
// 	pow := make([]int, 10)
// 	for i := range pow {
// 		pow[i] = 1 << uint(i) // == 2**i
// 	}
// 	for _, value := range pow {
// 		fmt.Printf("%d\n", value)
// 	}
// }