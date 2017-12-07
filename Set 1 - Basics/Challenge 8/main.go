// Detect 16 bit ECB encryption

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
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
	lineMap := make(map[int]int)

	for linenum, line := range fileContentArray {
		decodedHex, err := hex.DecodeString(string(line))
		check(err)
		lineMap[linenum] = ScoringSystem(decodedHex)
	}
	ecbline := getLowest(lineMap)
	fmt.Println(fileContentArray[ecbline])
}

func getLowest(lineMap map[int]int) int {
	lowestCount := 9999
	lowestLineNum := 0

	for linenum, count := range lineMap {
		if count < lowestCount {
			lowestCount = count
			lowestLineNum = linenum
		}
	}
	fmt.Println(lowestCount)
	return lowestLineNum
}

func ScoringSystem(bArray []byte) int {
	highFreq := make(map[byte]int)
	decodeLen := len(bArray)

	for i := 0; i < decodeLen; i++ {
		highFreq[bArray[i]] += 1
	}

	return len(highFreq)
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
