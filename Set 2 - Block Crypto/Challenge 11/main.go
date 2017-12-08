package main

import (
	"crypto/aes"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	data := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	executionTimes := 100
	encryptedData := []byte{}
	modeSelect := 9

	aesKeySize := 16
	ivKeySize := 16
	blockSize := 16

	for i := 0; i < executionTimes; i++ {
		aesKey := generateBytes(aesKeySize)
		ivKey := generateBytes(ivKeySize)

		modeSelect = rand.Intn(2)
		switch modeSelect {
		case 0:
			// ECB
			prependedData := prependBytes(data)
			appendedData := appendBytes(prependedData)
			encryptedData = encryptECB(appendedData, aesKey, blockSize)
		case 1:
			// CBC
			prependedData := prependBytes(data)
			appendedData := appendBytes(prependedData)
			encryptedData = encryptCBC(appendedData, aesKey, ivKey, blockSize)
		}
		// Mode: 0/1 ECB/CBC - Checker: true/false ECB/CBC
		fmt.Printf("Mode: %d Checker: %v\n", modeSelect, determineECB(encryptedData, blockSize))
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

func encryptCBC(data, key, iv []byte, blockSize int) []byte {
	finalResult := []byte{}

	if len(data)%blockSize != 0 {
		data = padMe(data, (blockSize - (len(data) % blockSize)))
	}

	for bs, be := 0, blockSize; bs < len(data); bs, be = bs+blockSize, be+blockSize {
		cstring := data[bs:be]
		encryptResult := cbcMode(cstring, key, iv, blockSize)
		finalResult = append(finalResult, encryptResult...)
		iv = encryptResult
	}
	return finalResult
}

func cbcMode(cstring, key, iv []byte, blockSize int) []byte {
	xordString := repeatingKeyXOR(cstring, iv)

	block, err := aes.NewCipher(key)
	check(err)

	encryptedData := make([]byte, len(xordString))
	block.Encrypt(encryptedData, xordString)

	return encryptedData
}

func determineECB(bArray []byte, blockSize int) bool {
	ecbMode := false

	blockSlices := [][]byte{}
	for bs, be := 0, blockSize; bs < len(bArray); bs, be = bs+blockSize, be+blockSize {
		blockSlices = append(blockSlices, bArray[bs:be])
	}

	decodeLen := len(blockSlices)
	for i := 0; i < decodeLen-1; i++ {
		for j := i + 1; j < decodeLen; j++ {
			if reflect.DeepEqual(blockSlices[i], blockSlices[j]) {
				ecbMode = true
				break
			}
		}
	}
	return ecbMode
}

func padMe(block []byte, paddingAmount int) []byte {
	for count := 0; count < paddingAmount; count++ {
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
		keyIterator++
		if keyIterator == keyCount {
			keyIterator = 0
		}
	}
	return resultArray
}

func generateBytes(keyLength int) []byte {
	key := make([]byte, keyLength)
	rand.Read(key)
	return key
}

func prependBytes(plaintext []byte) []byte {
	bytes := generateBytes((rand.Intn(5)) + 5)
	bytes = append(bytes, plaintext...)
	return bytes
}

func appendBytes(plaintext []byte) []byte {
	bytes := generateBytes((rand.Intn(5)) + 5)
	plaintext = append(plaintext, bytes...)
	return plaintext
}
