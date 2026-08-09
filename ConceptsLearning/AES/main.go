package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := []byte("12345678901234567890123456789012")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n",block)
	fmt.Println(block)
	fmt.Println(block.BlockSize())
}