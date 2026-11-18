package main

import (
	"fmt"
)

func main() {
	// CreatingMaps()
	// coreOperations()
	MissingKeyAndCommaOK()
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

func MissingKeyAndCommaOK() {
	m := map[string]int{"a": 2}
	fmt.Println(m["a"])   // 0   (key exists, value is 0)
	fmt.Println(m["zzz"]) // 0   (key missing, zero value)

	// v, ok := m["zzz"]
	// ok is true if the key exists, false otherwise
	// v is the value, or the zero value if missing

	if v, ok := m["a"]; ok { // okay is just a regular variable. you can name it anything
		fmt.Println("found", v, ok)
	}
}
