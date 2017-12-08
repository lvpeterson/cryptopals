package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	challengeFile = "encryptedData"
)

func main() {
	KEY := []byte("YELLOW SUBMARINE")
	IV := []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	finalResult := []byte{}
	blockSize := 16
	fileContents, err := ioutil.ReadFile(challengeFile)
	check(err)
	decodedContents, err := base64.StdEncoding.DecodeString(string(fileContents))
	check(err)

	for bs, be := 0, blockSize; bs < len(decodedContents); bs, be = bs+blockSize, be+blockSize {
		cstring := decodedContents[bs:be]
		decryptResult := cbcMode(cstring, KEY, IV, false, blockSize)
		finalResult = append(finalResult, decryptResult...)
		IV = decodedContents[bs:be]
	}

	fmt.Println(string(finalResult))

}

func cbcMode(cstring, key, iv []byte, mode bool, blockSize int) []byte {
	// mode: true = encrypt | false = decrypt
	// Order of Ops Encrypt: plaintext XOR with IV/Previous Cipherblock -> encrypt
	// Order of Ops Decrypt: ciphertext decrypt -> xor decrypted with IV/Previous cipherblock
	if mode {
		paddedCString := padMe(cstring, blockSize)
		xordString := repeatingKeyXOR(paddedCString, iv)
		encryptedString := encryptBlock(xordString, key)
		return encryptedString
	} else {
		decryptedString := decryptBlock(cstring, key)
		xordString := repeatingKeyXOR(decryptedString, iv)
		return xordString
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func encryptBlock(data, key []byte) []byte {

	block, err := aes.NewCipher(key)
	check(err)

	encryptedData := make([]byte, len(data))
	block.Encrypt(encryptedData, data)

	return encryptedData
}

func decryptBlock(data, key []byte) []byte {

	block, err := aes.NewCipher(key)
	check(err)

	decryptedData := make([]byte, len(data))
	block.Decrypt(decryptedData, data)

	return decryptedData
}

func padMe(block []byte, blockSize int) []byte {
	paddingLength := blockSize - len(block)
	for count := 0; count < paddingLength; count++ {
		block = append(block, '\x00')
	}
	return block
}

func repeatingKeyXOR(cstring, key []byte) []byte {
	resultArray := []byte{}

	stringCount := len(cstring)
	keyCount := len(key)
	keyIterator := 0

	for i := 0; i < stringCount; i++ {
		resultArray = append(resultArray, (cstring[i] ^ key[keyIterator]))
		keyIterator += 1
		if keyIterator == keyCount {
			keyIterator = 0
		}
	}
	return resultArray

}
