/*******************************************************************************************************

AES in ECB mode
The Base64-encoded content in this file (challenge7file.txt) has been encrypted via AES-128 in ECB mode under the key

"YELLOW SUBMARINE".
(case-sensitive, without the quotes; exactly 16 characters; I like "YELLOW SUBMARINE" because it's exactly 16 bytes long, and now you do too).

Decrypt it. You know the key, after all.

Easiest way: use OpenSSL::Cipher and give it AES-128-ECB as the cipher.

Do this with code.
You can obviously decrypt this using the OpenSSL command-line tool, but we're having you get ECB working in code for a reason. You'll need it a lot later on, and not just for attacking ECB.

*******************************************************************************************************/

package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"encoding/base64"
	"crypto/aes"
)

const (
	challengefile = "challenge7file.txt"

)

func main() {
	// Challenge Setup
	key := []byte("YELLOW SUBMARINE")

	fileContents, err := ioutil.ReadFile(challengefile)
	if err != nil {
        log.Fatal(err)
    }
    decodedContents, derr := base64.StdEncoding.DecodeString(string(fileContents))
    if derr != nil {
        log.Fatal(err)
    }

    decryptedData := decryptAes128ECB(decodedContents, key)
    fmt.Println (string(decryptedData))
}

func decryptAes128ECB(data, key []byte) []byte{

    blockSize := 16
    block, err := aes.NewCipher(key)
    if err != nil {
        log.Fatal(err)
    }

    decryptedData := make([]byte, len(data))
    for begChunk, endChunk := 0, blockSize; begChunk < len(data); begChunk, endChunk = begChunk+blockSize, endChunk+blockSize {
        block.Decrypt(decryptedData[begChunk:endChunk], data[begChunk:endChunk])
    }

    return decryptedData
}