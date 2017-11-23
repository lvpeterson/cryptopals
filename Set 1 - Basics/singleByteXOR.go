/*******************************************************************************************************

Single-byte XOR cipher
The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.

Achievement Unlocked
You now have our permission to make "ETAOIN SHRDLU" jokes on Twitter.

Secret Message: Cooking MC's like a pound of bacon

*******************************************************************************************************/

package main

import (
	"fmt"
	"encoding/hex"
)

func main(){

	// Challenge Setup
	challengeString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decodedHexCString, _ := hex.DecodeString(challengeString)
	fullCharSet := []byte("~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	charScore := make(map[byte]int)

	// Loop through each character and XOR
	charCount := len(fullCharSet)
	for i := 0; i < charCount; i++ {
		resultXOR := XORSingleBits(decodedHexCString, fullCharSet[i]);
		resultScore := ScoringSystem(resultXOR)
		charScore[fullCharSet[i]] = resultScore
	}

	byteKey := getBestCharScore(charScore)
	finalResult := XORSingleBits(decodedHexCString, byteKey)
	fmt.Println(string(finalResult))

}

// Single Bit XOR
func XORSingleBits(bArray1 []byte, character byte) []byte{
	// Assumes both arrays are same length
	resultArray := []byte{}
	count := len(bArray1)
	for i := 0;  i < count; i++{
		resultArray = append(resultArray, (bArray1[i] ^ character))
	}
	return resultArray
}

// Scoring system for most likely to be used
func ScoringSystem(bArray []byte) int{
	highFreq := []byte("etaoinshrdluETAOINSHRDLU ")
	highFreqLen := len(highFreq)
	decodeLen := len(bArray)
	score := 0

	for i := 0; i < decodeLen; i++{
		for j := 0; j < highFreqLen; j++{
			if bArray[i] == highFreq[j] {
				score += 1
				break
			}
		}
	}
	return score

}

func getBestCharScore(charScore map[byte]int) byte{
	topScore := 0
	topByteScore := byte('0') 

	for byteVal, charScoreNum := range charScore {
		if (charScoreNum > topScore){
			topScore = charScoreNum	
			topByteScore = byteVal
		}
	}
	return topByteScore
}
