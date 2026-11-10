package main

import (
	"fmt"
)

func main() {
	// Slice()
	// DeleteWithOrderPreservation()
	DeleteWithNoOrderPresevation()
}

func Slice() {
	slice := []int{100, 400, 200, 300, 500}

	preSlicer := slice[0:]
	preSlicer2 := slice[:1]

	fmt.Println(preSlicer)
	fmt.Println(preSlicer2)
}

func lastElement() {
	slice := []int{100, 400, 200, 300, 500}

	last := slice[len(slice)-1]

	fmt.Println(last)
}

func DeleteWithOrderPreservation() {
	slice := []int{100, 400, 200, 300, 500}

	fmt.Println(slice[:4])
	fmt.Println(slice[3:])
	fmt.Println(slice[1+1:])
	delete200 := append(slice[:1], slice[1+1:]...)
	fmt.Println(delete200)
}

func DeleteWithNoOrderPresevation() {
	s := []int{100, 400, 200, 300, 500}
	fmt.Println(s[len(s)-1])
	fmt.Println(s[:len(s)-1])


	s[2] = s[len(s)-1]
	fmt.Println(s[3])
	fmt.Println(s)
	s = s[:len(s)-1]

	fmt.Println(s)
}
