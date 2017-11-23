/*******************************************************************************************************

Break repeating-key XOR
It is officially on, now.
This challenge isn't conceptually hard, but it involves actual error-prone coding. 
The other challenges in this set are there to bring you up to speed. This one is there to qualify you. If you can do this one, you're probably just fine up to Set 6.

There's a file here (challenge6file.txt). It's been base64'd after being encrypted with repeating-key XOR.

Decrypt it.

Here's how:

Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
Write a function to compute the edit distance/Hamming distance between two strings. The Hamming distance is just the number of differing bits. The distance between:
this is a test
and
wokka wokka!!!
is 37. Make sure your code agrees before you proceed.
For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes, and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
The KEYSIZE with the smallest normalized edit distance is probably the key. You could proceed perhaps with the smallest 2-3 KEYSIZE values. 
Or take 4 KEYSIZE blocks instead of 2 and average the distances.
Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
Now transpose the blocks: make a block that is the first byte of every block, and a block that is the second byte of every block, and so on.
Solve each block as if it was single-character XOR. You already have code to do this.
For each block, the single-byte XOR key that produces the best looking histogram is the repeating-key XOR key byte for that block. Put them together and you have the key.
This code is going to turn out to be surprisingly useful later on. Breaking repeating-key XOR ("Vigenere") statistically is obviously an academic exercise, a "Crypto 101" thing. 
But more people "know how" to break it than can actually break it, and a similar technique breaks something much more important.

No, that's not a mistake.
We get more tech support questions for this challenge than any of the other ones. We promise, there aren't any blatant errors in this text. 
In particular: the "wokka wokka!!!" edit distance really is 37.

Not Padded:
-----------
C1: [1110100 1101000 1101001 1110011 100000 1101001 1110011 100000 1100001 100000 1110100 1100101 1110011 1110100]
C2: [1110111 1101111 1101011 1101011 1100001 100000 1110111 1101111 1101011 1101011 1100001 100001 100001 100001]

Padded:
-------
C1: [1110100 1101000 1101001 1110011 0100000 1101001 1110011 0100000 1100001 0100000 1110100 1100101 1110011 1110100]
C2: [1110111 1101111 1101011 1101011 1100001 0100000 1110111 1101111 1101011 1101011 1100001 0100001 0100001 0100001]
     0000011 0000111 0000010 0011000 1000001 1001001 0000100 1001111 0001010 1001011 0010101 1000100 1010010 1010101
		2		3		1		2		2		3		1		5		2		4		3		2		3		4		= 37

Key: Terminator X: Bring the noise

*******************************************************************************************************/

package main

import(
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"sort"
	"encoding/base64"
)

const (
	MAXKEYSIZE = 40
	MINKEYSIZE = 2
	challengefile = "challenge6file.txt"
)

func main() {

	fileContents, err := ioutil.ReadFile(challengefile)
	if err != nil {
        log.Fatal(err)
    }
    decodedContents, derr := base64.StdEncoding.DecodeString(string(fileContents))
    if derr != nil {
        log.Fatal(err)
    }
    //fmt.Println(fileContents)
    keySize := findKeySize(decodedContents)

	blocks := createBlocks(keySize, decodedContents)

	transposeBlocks := transposeBlocks(blocks)

	// Solve Encryption
	//-----------------------------------------------------------------
	key := getKey(transposeBlocks)
	fmt.Printf("Key: %s\n\n",string(key))
	finalString := breakRepeatingXOR(key, decodedContents)
	fmt.Println(string(finalString))
}

func breakRepeatingXOR(key, encryptedText []byte) []byte{
	resultArray := []byte{}

	stringCount := len(encryptedText)
	keyCount := len(key)
	keyIterator := 0

	for i := 0;  i < stringCount; i++{
		resultArray = append(resultArray, (encryptedText[i] ^ key[keyIterator]))
		keyIterator += 1
		if keyIterator == keyCount{
			keyIterator = 0
		}
	}
	return resultArray
}

// Get Repeating Key Used
func getKey(transposeBlocks map[int][]byte) []byte{
	fullCharSet := []byte("~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	key := []byte{}
	

	challengeLen := len(transposeBlocks) //challengeLines
	// Get best scores from each line
	
	for i := 0; i < challengeLen; i++{
		charScore := make(map[byte]int)

		// Loop through each character and XOR
		charCount := len(fullCharSet)
		for j := 0; j < charCount; j++ {
			resultXOR := XORSingleBits(transposeBlocks[i], fullCharSet[j]);
			resultScore := ScoringSystem(resultXOR)
			charScore[fullCharSet[j]] = resultScore
		}

		byteKey := getBestCharScore(charScore)
		key = append(key, byteKey)
	}

	return key
}

// Convert String to Array of Chars
func stringToArray(fString string) []string {
	return strings.Split(fString, "")
}

// Create Blocks of KEYSIZE
func createBlocks(keyLength int, encryptedText []byte) map[int][]byte{
	blocks := make(map[int][]byte)
	count := 0

	for currentLocation := 0; currentLocation < len(encryptedText); currentLocation += keyLength {
		blocks[count] = encryptedText[currentLocation:currentLocation+keyLength]
		count ++	
	}

	return blocks
}

// Transpose Created Blocks
func transposeBlocks(blocks map[int][]byte) map[int][]byte{
	iterations := len(blocks[0]) // sets number of transposes to do
	transposeBlocks := make(map[int][]byte)

	for i := 0; i < iterations; i++ {
		transpose := []byte{}
		// do some stuff
		for key := 0; key < len(blocks); key++ {
			byteArray := blocks[key]
			transpose = append(transpose, byteArray[i])
		}
		transposeBlocks[i] = transpose
	}
	return transposeBlocks
}

// find KEYSIZE for encrypted text - Could be wrong
func findKeySize(encryptedText []byte) int{

	normalizedKeys := make(map[int]float64)

	for currentKeySize := MINKEYSIZE; currentKeySize < MAXKEYSIZE + 1; currentKeySize ++{

		// Get Distance and normalize by dividing by keylength
		encryptedSlice1 := encryptedText[:currentKeySize]
		encryptedSlice2 := encryptedText[currentKeySize:currentKeySize*2]
		encryptedSlice3 := encryptedText[currentKeySize*2:currentKeySize*3]
		encryptedSlice4 := encryptedText[currentKeySize*3:currentKeySize*4]

		hd1 := float64(bitwiseHammingDist(encryptedSlice1, encryptedSlice2)) / float64(currentKeySize)
		hd2 := float64(bitwiseHammingDist(encryptedSlice1, encryptedSlice3)) / float64(currentKeySize)
		hd3 := float64(bitwiseHammingDist(encryptedSlice1, encryptedSlice4)) / float64(currentKeySize)

		//normalizedDist := float64(hammingDist) / float64(currentKeySize)

		normalizedKeys[currentKeySize] = (hd1 + hd2 + hd3)/3
	} 

	keyLength := checkKeys(normalizedKeys, encryptedText)

	return keyLength
}

func testKeyLength(keyLength int, encryptedText []byte) []byte{
	blocks := createBlocks(keyLength, encryptedText)
	transposeBlocks := transposeBlocks(blocks)

	// Solve Encryption
	//-----------------------------------------------------------------
	key := getKey(transposeBlocks)
	return key
} 

func checkKeys(unsortedMap map[int]float64, encryptedText []byte) int {
	numOfChecks := 5
	currentChecks := 0

	testKeys := make(map[int][]byte)
	sortedMap := make(map[float64][]int)
    var sortedFloatArray []float64
    // sort map by values not keys
    for keySize, hammingNormDist := range unsortedMap {
            sortedMap[hammingNormDist] = append(sortedMap[hammingNormDist], keySize)
    }
    for hammingNormDist := range sortedMap {
            sortedFloatArray = append(sortedFloatArray, hammingNormDist)
    }
    sort.Sort(sort.Float64Slice(sortedFloatArray))
    for _, hammingNormDist := range sortedFloatArray {
            for _, keySize := range sortedMap[hammingNormDist] {
            		if currentChecks < numOfChecks {
            			testKeys[keySize] = testKeyLength(keySize, encryptedText)
            			currentChecks ++
            		}
                    
            }
    }
    // Get best scoring Key
    topScore := 0
    topKeyLength := 0
    for keySize, key := range testKeys {
		score := ScoringSystem(key)
		if topScore < score{
			topScore = score
			topKeyLength = keySize
		}
    }

    return topKeyLength
}

// Perform bitwise Hamming Dist Check
func bitwiseHammingDist(fString, sString []byte) int{
	var count int
	stringLength := len(fString)
	for i := 0; i < stringLength; i++ {

		// Get Binary Stuff Ready
		fStringByte := fmt.Sprintf("%b",fString[i])
		sStringByte := fmt.Sprintf("%b",sString[i])

		// Pad the binary
		paddedfStringByte := padBinary(fStringByte)
		paddedsStringByte := padBinary(sStringByte)

		// Convert to array
		paddedfArray := stringToArray(paddedfStringByte)
		paddedsArray := stringToArray(paddedsStringByte)

		for j := 0; j < len(paddedfArray); j++ {
			if string(paddedfArray[j]) != string(paddedsArray[j]) {
				count++
			}
		}	
	}
	return count
}

// Get Lowest Score for KEYSIZE
func getLowestFloatScoreKeySize(keyNormalized map[int]float64) int{
	lowestScore := float64(99999999) // Large Number to compare to
	finalKeySize := 0

	for keySize, scoreNum := range keyNormalized {
		fmt.Printf("Key: %d Score: %f\n", keySize, scoreNum)
		if (scoreNum < lowestScore){
			lowestScore = scoreNum
			finalKeySize = keySize	
		}
	}
	return finalKeySize
}

// Left Pad Binay Number
func padBinary(binaryNum string) string {
	// left pad with 0's to make it length of 8
	binLength := len(binaryNum)
	padCount := 8
	pad := "0"

	for i := binLength; i < padCount; i++{
        binaryNum = pad + binaryNum
    }

    return binaryNum
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