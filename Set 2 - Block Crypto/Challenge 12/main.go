package main

import (
	"fmt"

	cryptoHelp "github.com/lvpeterson/cryptopals/cryptoHelp"
)

func main() {
	aesKeySize := 16
	aesKey := cryptoHelp.generateBytes(aesKeySize)
	fmt.Println(aesKey)
}
