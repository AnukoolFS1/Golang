package main

import (
	"fmt"
)

func main() {
	CreatingMaps()
	coreOperations()
}

func CreatingMaps() {
	m1 := make(map[string]int)           // empty, ready to use
	m2 := map[string]int{"a": 1, "b": 2} // literal with data
	m3 := make(map[string]int, 100)      // size hint: avoids early regrowth
	var m4 map[string]int                // nil map (no storage yet)

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
}

func coreOperations() {
	m := make(map[string]int)

	m["apple"] = 5  // insert or overwrite
	x := m["apple"] // read -> 5
	n := len(m)     // number of entries

	fmt.Println(m)
	fmt.Println(x)
	fmt.Println(n)

	delete(m, "apple") // remove (deleting a missing key is a safe no-op)

	fmt.Println("After Delete")
	fmt.Println(m)
	fmt.Println(x)
	fmt.Println(n)
}
