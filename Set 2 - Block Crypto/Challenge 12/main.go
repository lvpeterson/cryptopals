package main

import (
	"encoding/base64"
	"fmt"

	"github.com/lvpeterson/cryptopals/crypt"
)

func main() {
	aesKeySize := 16
	aesKey := crypt.GenerateBytes(aesKeySize)

	b64string := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
		"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
		"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
		"YnkK"
	myString := []byte("AAAAAAAAAAAAAAABBBBBBBBBBBBBCCCCCCCCCCCCCCCC")
	decodedString, err := base64.StdEncoding.DecodeString(b64string)
	crypt.Check(err)

	combinedStr := append(myString, decodedString...)
	fmt.Println(combinedStr)

	fmt.Println(aesKey)
}
