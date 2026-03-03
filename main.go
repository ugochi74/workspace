package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run . <input.txt> <output.txt>")
        return
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    if err := run(inputFile, outputFile); err != nil {
        fmt.Println("Error:", err)
    }
}
