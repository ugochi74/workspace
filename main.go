package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text, err := ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := Process(text)

	err = WriteFile(outputFile, result)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Processing completed successfully.")
}