// Detect 16 bit ECB encryption

package main

import(
	"fmt"
	"log"
	"bufio"
	"os"
	"encoding/hex"
	//"crypto/aes"
)

const (
	challengefile = "encryptedData"

)

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func main() {
	fileContentArray, err := fileToArray(challengefile)
	check(err)
	fmt.Println(len(fileContentArray))

	for _,line := range fileContentArray {
	    decodedHex, err := hex.DecodeString(string(line))
	    check(err)
	    ScoringSystem(decodedHex)
	}

}

func ScoringSystem(bArray []byte) int{
	highFreq := make(map[byte]int)
	decodeLen := len(bArray)
	score := 0

	for i := 0; i < decodeLen; i++{
		highFreq[bArray[i]] += 1
	}

	fmt.Println(highFreq)
	return score
}

func fileToArray (filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close ()

    lines := []string{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines, scanner.Err()
}