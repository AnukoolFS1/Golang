package main

import (
	"fmt"
	"strings"
)

func countWord(str string) string {
	wordSlice := strings.Split(str, " ")

	fmt.Println(wordSlice)

	var wordCountMap map[string]int = make(map[string]int)

	for _, word := range wordSlice {
		if wordCountMap[word] == 0 {
			wordCountMap[word] = 1
		} else {
			wordCountMap[word] += 1
		}
	}

	for word, count := range wordCountMap {
		fmt.Println(word, count)
	}

	return ""
}

func main() {
	countWord("Hello World Hello Anukool")
}
