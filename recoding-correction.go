package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func hexToDecimal(hexStr string) (int64, error) {
	dec, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return 0, err
	}
	return dec, nil
}

func binToDecimal(binStr string) (int64, error) {
	dec, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return 0, err
	}
	return dec, nil
}

func capitalizeWord(word string) string {
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])

	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func toLower(word string) string {
	return strings.ToLower(word)
}

func toUpper(word string) string {
	return strings.ToUpper(word)
}
func main() {
	fmt.Println(hexToDecimal("1E")) // prints 30
	fmt.Println(hexToDecimal("FF")) // prints 255
	fmt.Println(binToDecimal("10"))
	fmt.Println(binToDecimal("1010"))
	fmt.Println(binToDecimal("11111111"))
	fmt.Println(capitalizeWord("hELLO"))
	fmt.Println(capitalizeWord("WORLD"))
	fmt.Println(toLower("hELLO"))
	fmt.Println(toUpper("world"))
}
