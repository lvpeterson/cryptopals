package main

import (
	"crypto/aes"
	"fmt"
	"log"
	"math/rand"
)

func main() {
	aesKeySize := 16
	aesKey := generateBytes(aesKeySize)
	fmt.Println(aesKey)
}

func encryptECB(data, key []byte, blockSize int) []byte {
	if len(data)%blockSize != 0 {
		data = padMe(data, (blockSize - (len(data) % blockSize)))
	}
	block, err := aes.NewCipher(key)
	check(err)

	encryptedData := make([]byte, len(data))
	for bs, be := 0, blockSize; bs < len(encryptedData); bs, be = bs+blockSize, be+blockSize {
		block.Encrypt(encryptedData[bs:be], data[bs:be])
	}
	return encryptedData
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func padMe(block []byte, paddingAmount int) []byte {
	for count := 0; count < paddingAmount; count++ {
		block = append(block, '\x00')
	}
	return block
}

func generateBytes(keyLength int) []byte {
	key := make([]byte, keyLength)
	rand.Read(key)
	return key
}
