/*******************************************************************************************************

Detect single-character XOR
One of the 60-character strings in this file (challenge4file.txt) has been encrypted by single-character XOR.

Find it.

(Your code from #3 should help.)

Secret Message: Now that the party is jumping

*******************************************************************************************************/

package main

import (
	"fmt"
	"encoding/hex"
	"os"
	"bufio"
)

func main(){

	// Challenge Setup
	// Load line from file -
	challengeLines, err := fileToArray("challenge4file.txt")
	if err != nil{
		fmt.Println(err)
	}

	fullCharSet := []byte("~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	stringScore := make(map[string]byte)

	challengeLen := len(challengeLines)
	// Get best scores from each line
	for i := 0; i < challengeLen; i++{
		decodedHexCString, _ := hex.DecodeString(challengeLines[i])
		charScore := make(map[byte]int)

		// Loop through each character and XOR
		charCount := len(fullCharSet)
		for i := 0; i < charCount; i++ {
			resultXOR := XORSingleBits(decodedHexCString, fullCharSet[i]);
			resultScore := ScoringSystem(resultXOR)
			charScore[fullCharSet[i]] = resultScore
		}

		byteKey := getBestCharScore(charScore)
		stringScore[challengeLines[i]] = byteKey
	}

	finalStringScore := make(map[string]int)
	// Get highest score from the previous step
	for hexString, bestByte := range stringScore {
		decodedHexString, _:= hex.DecodeString(hexString)
		finalXOR := XORSingleBits(decodedHexString, bestByte)
		finalScore := ScoringSystem(finalXOR)
		finalStringScore[string(finalXOR)] = finalScore
	}

	finalString := getBestStringScore(finalStringScore)
	fmt.Println(finalString)
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

// Get most likely char that decrypts message
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

// Get most likely string that is correctly decoded
func getBestStringScore(stringScore map[string]int) string{
	topScore := 0
	topStringScore := ""

	for stringVal, scoreNum := range stringScore {
		if (scoreNum > topScore){
			topScore = scoreNum	
			topStringScore = stringVal
		}
	}
	return topStringScore
}

// Take file input and put into array line by line
func fileToArray (filePath string) ([]string, error) {
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
