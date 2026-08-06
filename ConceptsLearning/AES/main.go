package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := []byte("1234567890123456789012345678901")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n",block)
	fmt.Println(block)
	fmt.Println(block.BlockSize())
}