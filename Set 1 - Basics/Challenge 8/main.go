// Detect 16 bit ECB encryption

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"reflect"
	//"crypto/aes"
)

const (
	challengefile = "encryptedData"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileContentArray, err := fileToArray(challengefile)
	check(err)

	blockSize := 16

	for linenum, line := range fileContentArray {
		decodedHex, err := hex.DecodeString(string(line))
		check(err)
		if determineECB(decodedHex, blockSize) {
			fmt.Printf("ECB Found at line: %d with String: %s", linenum, line)
		}
	}
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

func fileToArray(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
