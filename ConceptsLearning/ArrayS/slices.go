package main

import "fmt"

// defining array in GO
// var variableName = [length]type{el, el1, el2...} /// type of this is [l]T

// func testingSlice() {
// 	slicedata := []int{1, 2, 3, 4, 5, 6}
// 	slicePointer := &slicedata

// 	fmt.Printf("slicePointer %p\n", slicePointer)

// 	fmt.Printf("slicePointer %p\n", slicePointer)
// }

func main() {

	var arr [5]int = [5]int{1, 2, 3, 4, 5} // cap stand for capacity

	var slicedArr1 = arr[1:4]
	var slicedArr2 = arr[5:]

	fmt.Printf("SlicedArr %v\n", slicedArr1)
	fmt.Printf("%v\n", cap(slicedArr1))
	fmt.Printf("%v\n\n", len(slicedArr1))

	fmt.Printf("slicesArr2 %v\n", slicedArr2)
	fmt.Printf("%v\n", cap(slicedArr2))
	fmt.Printf("%v\n\n", len(slicedArr2))

	slicedArr2 = append(slicedArr2, 6)

	fmt.Printf("-%v\n", slicedArr2)
	fmt.Printf("%v\n", cap(slicedArr2))
	fmt.Printf("%v\n\n", len(slicedArr2))

	fmt.Printf("arr %v\n", arr)
	fmt.Printf("%v\n", cap(arr))
	fmt.Printf("%v\n", len(arr))
	// testingSlice()
	// SlicesOfSlice()
	// fmt.Println(nil)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// fmt.Println(s)

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
