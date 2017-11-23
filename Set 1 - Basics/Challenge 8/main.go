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
	count := 0
	for _,line := range fileContentArray {
	    decodedHex, err := hex.DecodeString(string(line))
	    check(err)
	    if(detectECB(decodedHex)){
	    	count ++
	    }
	}



    fmt.Println(count)

}

func detectECB(data []byte) bool {
	blockSize := 16
	if (len(data) % blockSize == 0){
		return true
	}

	return false
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