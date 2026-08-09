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

/*
Security Hashes

For passwords, we use cryptographic hash functions like:

SHA-256
SHA-512
bcrypt
Argon2
scrypt

These are designed with special security properties.
*/