package main

import (
	"fmt"
	"strings"
)

func countWord(str string) { //Often it's better for a function to return data and let the caller decide what to do with it.

	// Before the Split, it could have been better to lowercase the words first to avoid duplication of words
	// str = strings.ToLower(str)
	wordSlice := strings.Split(str, " ") // Fields works better with whitespaces
	// action => strings.Fields("Hello   World") notice the space
	// result => []string{"Hello", "World"}

	fmt.Println(wordSlice)

	var wordCountMap map[string]int = make(map[string]int)

	for _, word := range wordSlice {
		// Since initially every value is 0 we could have just done the
		// => wordCountMap[word]++
		if wordCountMap[word] == 0 {
			wordCountMap[word] = 1
		} else {
			wordCountMap[word] += 1
		}
	}

	for word, count := range wordCountMap {
		fmt.Println(word, count)
	}
	// fmt.Println(wordCountMap)

}

func main() {
	countWord("Hello World Hello Anukool")
}
