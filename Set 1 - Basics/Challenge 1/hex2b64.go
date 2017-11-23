/*******************************************************************************************************

Convert hex to base64
The string:

49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
Should produce:

SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
So go ahead and make that happen. You'll need to use this code for the rest of the exercises.

Cryptopals Rule
Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

Final Message: I'm killing your brain like a poisonous mushroom

*******************************************************************************************************/

package main

import (
	"fmt"
	"encoding/hex"
	"encoding/base64"
)

func main(){

	hexstring := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	decodedHex, _ := hex.DecodeString(hexstring)
	b64Encoded := base64.StdEncoding.EncodeToString(decodedHex)
	fmt.Println(b64Encoded)

}