package main

import (
	"fmt"

	"github.com/lvpeterson/cryptopals/crypt"
)

func main() {
	aesKeySize := 16
	aesKey := crypt.GenerateBytes(aesKeySize)
	fmt.Println(aesKey)

}
