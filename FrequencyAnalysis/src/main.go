package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Format: go run cipher-text.file")
		os.Exit(-1)
	}

	// Get cipher text file and frequency files
	cipherTextFile := os.Args[1]

	// Read cipher text file and profile frequency of characters
	charFreqMap, err := readFileAndCharFreq(cipherTextFile)
	if err != nil {
		fmt.Println("Something went wrong while calculating frequency: ", err)
		os.Exit(-1)
	}

	fmt.Println("Frequency of characters in cipher text is")
	printFrequency(charFreqMap)
}

/**
 * Function to read the file and count number of times each character appears
 * in the string. Find out frequency of occurance based on the total number of
 * characters read.
 */
func readFileAndCharFreq(file string) (map[byte]float32, error) {

	// read file and return it as string
	fileAsByteArr, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	fileAsStringTmp := string(fileAsByteArr)
	fileAsString := strings.TrimSuffix(strings.ToLower(fileAsStringTmp), "\n")

	// count each character occurrences
	charCount := make(map[byte]int)
	totalChars := 0
	// initialize all with 0
	for index := 0; index < 26; index++ {
		character := 'a' + byte(index)
		charCount[character] = 0
	}
	for _, character := range fileAsString {
		if string(character) == " " || string(character) == "\n" {
			continue
		}
		charCount[byte(character)] = charCount[byte(character)] + 1
		totalChars++
	}

	// Find out frequency
	charFreq := make(map[byte]float32)
	for character, count := range charCount {
		charFreq[character] = float32(count) / float32(totalChars)
	}
	return charFreq, nil
}

/**
 * Function to print frequency map
 */
func printFrequency(freqMap map[byte]float32) {
	for character, freq := range freqMap {
		fmt.Println(string(character), " : ", freq)
	}
}
