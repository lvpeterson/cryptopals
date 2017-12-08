package main

import (
	"fmt"

	"github.com/lvpeterson/cryptopals/crypt"
)

func main() {
	aesKeySize := 16
	aesKey := crypt.generateBytes(aesKeySize)
	fmt.Println(aesKey)

}
