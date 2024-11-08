package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func rot1(r rune) rune {
	// Apply ROT1 transformation for all ASCII characters
	return (r + 1) % 128
}

func rot1Decrypt(r rune) rune {
	// Reverse ROT1 transformation for all ASCII characters
	return (r - 1 + 128) % 128
}

func processFile(inputFile string, transformFunc func(rune) rune) (string, error) {
	// Read the input file
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	// Apply the transformation
	transformed := []rune{}
	for _, r := range string(data) {
		transformed = append(transformed, transformFunc(r))
	}

	return string(transformed), nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run rot1.go <encrypt|decrypt> <input_file>")
		return
	}

	action := os.Args[1]
	inputFile := os.Args[2]

	var transformFunc func(rune) rune

	// Determine whether to encrypt or decrypt
	if action == "encrypt" {
		transformFunc = rot1
	} else if action == "decrypt" {
		transformFunc = rot1Decrypt
	} else {
		fmt.Println("Invalid action. Use 'encrypt' or 'decrypt'.")
		return
	}

	// Process the file and output the result to the terminal
	result, err := processFile(inputFile, transformFunc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
