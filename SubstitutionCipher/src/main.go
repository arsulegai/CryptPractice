package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run cipher-text.file substitution.file")
		os.Exit(-1)
	}

	cipherTextFile := os.Args[1]
	substitutionFile := os.Args[2]

	// read substitution file, example of this file would be
	// a x
	// b t
	// ..
	// where a is replaces with x, b with t
	// TODO: Improve input validation
	subMap, err := readSubstitutionFile(substitutionFile)
	if err != nil {
		fmt.Println("Unexpected error : ", err)
		os.Exit(-1)
	}

	// read the cipher text
	cipherTextByteArr, err := ioutil.ReadFile(cipherTextFile)
	if err != nil {
		fmt.Println("Unexpected error : ", err)
		os.Exit(-1)
	}
	cipherTextString := string(cipherTextByteArr)

	// substitute as per the input
	plainText := substitute(cipherTextString, subMap)
	// print the plainText
	printTextFile(plainText)
}

/**
 * This function reads substitution file
 */
func readSubstitutionFile(file string) (map[byte]byte, error) {

	fileHandler, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fileHandler.Close()

	subsTable := make(map[byte]byte)

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		line := scanner.Text()
		subsTmp := strings.SplitN(line, " ", 2)
		subsTable[byte(subsTmp[0][0])] = byte(subsTmp[1][0])
	}
	return subsTable, nil
}

/**
 * This function substitutes the given string using the map
 */
func substitute(cipher string, subMap map[byte]byte) string {

	plainText := []rune(cipher)
	for index, character := range cipher {
		value, ok := subMap[byte(character)]
		if ok {
			plainText[index] = rune(value)
		}
	}
	return string(plainText)
}

/**
 * This function prints the deciphered plaintext
 */
func printTextFile(text string) {
	fmt.Println(text)
}
