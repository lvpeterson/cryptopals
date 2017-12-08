package crypt

import (
	"crypto/aes"
	"crypto/rand"
	"log"
	"reflect"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// --------------------------------------------------------------------
// Encryption Modes:
// --------------------------------------------------------------------

// Encrypt ECB Mode
// --------------------------------------------------------------------
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

// Decrypt ECB Mode
// --------------------------------------------------------------------
func decryptECB(data, key []byte, blockSize int) []byte {
	if len(data)%blockSize != 0 {
		data = padMe(data, (blockSize - (len(data) % blockSize)))
	}
	block, err := aes.NewCipher(key)
	check(err)

	encryptedData := make([]byte, len(data))
	for bs, be := 0, blockSize; bs < len(encryptedData); bs, be = bs+blockSize, be+blockSize {
		block.Decrypt(encryptedData[bs:be], data[bs:be])
	}
	return encryptedData
}

// Encrypt CBC Mode
// --------------------------------------------------------------------
func encryptCBC(data, key, iv []byte, blockSize int) []byte {
	finalResult := []byte{}

	if len(data)%blockSize != 0 {
		data = padMe(data, (blockSize - (len(data) % blockSize)))
	}

	for bs, be := 0, blockSize; bs < len(data); bs, be = bs+blockSize, be+blockSize {
		cstring := data[bs:be]
		xordString := repeatingKeyXOR(cstring, iv)

		block, err := aes.NewCipher(key)
		check(err)

		encryptedData := make([]byte, len(xordString))
		block.Encrypt(encryptedData, xordString)

		finalResult = append(finalResult, encryptedData...)
		iv = encryptedData
	}
	return finalResult
}

// Decrypt CBC Mode
// --------------------------------------------------------------------
func decryptCBC(data, key, iv []byte, blockSize int) []byte {
	finalResult := []byte{}

	if len(data)%blockSize != 0 {
		data = padMe(data, (blockSize - (len(data) % blockSize)))
	}

	for bs, be := 0, blockSize; bs < len(data); bs, be = bs+blockSize, be+blockSize {
		cstring := data[bs:be]

		block, err := aes.NewCipher(key)
		check(err)

		decryptedData := make([]byte, len(cstring))
		block.Decrypt(decryptedData, cstring)

		xordString := repeatingKeyXOR(decryptedData, iv)
		finalResult = append(finalResult, xordString...)
		iv = cstring
	}
	return finalResult
}

// --------------------------------------------------------------------
// XOR Functions
// --------------------------------------------------------------------

// repeatingKeyXOR
// --------------------------------------------------------------------
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

// --------------------------------------------------------------------
// Misc. Functions
// --------------------------------------------------------------------

// Padding Function
// --------------------------------------------------------------------
func padMe(block []byte, paddingAmount int) []byte {
	for count := 0; count < paddingAmount; count++ {
		block = append(block, '\x00')
	}
	return block
}

// Generate X bytes in byte array
// --------------------------------------------------------------------
func generateBytes(keyLength int) []byte {
	key := make([]byte, keyLength)
	rand.Read(key)
	return key
}

// Determine Cipher Block
// --------------------------------------------------------------------
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
