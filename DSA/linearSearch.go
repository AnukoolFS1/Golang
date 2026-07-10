package main

func LinearSearch(list []int, search int) int {
	for i, n := range list {
		if n == search {
			return i
		}
	}

	return -1
}
