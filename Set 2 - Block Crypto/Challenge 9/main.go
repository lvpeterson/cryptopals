package main

import (
	"fmt"
	"log"
)

func main() {
	paddingString := []byte("hello")
	padMe(paddingString, 8)
	padMe(paddingString, 16)
}

func padMe(block []byte, blockSize int) {
	paddingLength := blockSize - len(block)
	for count := 0; count < paddingLength; count++ {
		block = append(block, '\x00')
	}
	fmt.Println(block)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
