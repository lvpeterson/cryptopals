/*******************************************************************************************************

Fixed XOR
Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c
... after hex decoding, and when XOR'd against:

686974207468652062756c6c277320657965
... should produce:

746865206b696420646f6e277420706c6179

Final Message: the kid don't play

*******************************************************************************************************/

package main

import (
	"fmt"
	"encoding/hex"
)

func main(){

	// Challenge Setup
	firstString := "1c0111001f010100061a024b53535009181c"
	secondString := "686974207468652062756c6c277320657965"
	decodedHexS1, _ := hex.DecodeString(firstString)
	decodedHexS2, _ := hex.DecodeString(secondString)

	hexXOR := XORbits(decodedHexS1, decodedHexS2)
	
	finalResult := hex.EncodeToString(hexXOR)
	fmt.Println(finalResult)

	// Print Secret Message
	fmt.Println(string(hexXOR))
}

// Bitwise XOR each byte in array
func XORbits(bArray1, bArray2 []byte) []byte{
	// Assumes both arrays are same length
	resultArray := []byte{}
	count := len(bArray1)
	for i := 0;  i < count; i++{
		resultArray = append(resultArray, (bArray1[i] ^ bArray2[i]))
	}

	return resultArray
}