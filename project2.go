package internal

import (
	"fmt"
	"os"
)

// ReadFile reads the content of a file and returns it as a string
func ReadFile(filename string) string {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return string(data)
}

```go
writer.go


package utils

import "testing"

func TestHexToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1E", 30},
		{"FF", 255},
		{"0", 0},
		{"A", 10},
	}

	for _, test := range tests {
		result := HexToDecimal(test.input)
		if result != test.expected {
			t.Errorf("HexToDecimal(%s) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestBinToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"10", 2},
		{"1010", 10},
		{"0", 0},
		{"1", 1},
	}

	for _, test := range tests {
		result := BinToDecimal(test.input)
		if result != test.expected {
			t.Errorf("BinToDecimal(%s) = %d; want %d", test.input, result, test.expected)
		}
	}
}
```

```go
for learning

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}

	first := strings.ToUpper(string(word[0]))
	rest := strings.ToLower(word[1:])

	return first + rest
}
```

```go
modifiers


package internal

import (
	"strconv"
	"strings"
)

// ApplyModifiers processes hex, bin, up, low, cap
func ApplyModifiers(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		switch words[i] {

		case "(hex)":
			if i > 0 {
				num, _ := strconv.ParseInt(words[i-1], 16, 64)
				words[i-1] = strconv.Itoa(int(num))
			}
			words = removeWord(words, i)
			i--

		case "(bin)":
			if i > 0 {
				num, _ := strconv.ParseInt(words[i-1], 2, 64)
				words[i-1] = strconv.Itoa(int(num))
			}
			words = removeWord(words, i)
			i--

		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = removeWord(words, i)
			i--

		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
			words = removeWord(words, i)
			i--

		case "(cap)":
			if i > 0 {
				words[i-1] = capitalize(words[i-1])
			}
			words = removeWord(words, i)
			i--
		}
	}

	return strings.Join(words, " ")
}
```

```go
punctuation

package internal

import "strings"

// FixPunctuation fixes spacing around punctuation marks
func FixPunctuation(text string) string {

	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuations {

		// remove space before punctuation
		text = strings.ReplaceAll(text, " "+p, p)

		// ensure space after punctuation
		text = strings.ReplaceAll(text, p+" ", p+" ")
	}

	return text
}
```

```go
quotes

package internal

import "strings"

// FixQuotes fixes spaces inside single quotes
func FixQuotes(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		// if a word starts with a quote
		if words[i] == "'" && i+2 < len(words) {

			// combine quote + word
			words[i] = "'" + words[i+1]

			// combine word + quote
			words[i] = words[i] + "'"

			// remove used words
			words = append(words[:i+1], words[i+3:]...)
		}
	}

	return strings.Join(words, " ")
}
```
```go
another example of quotes


package internal

import "strings"

func FixQuotes(text string) string {

	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")

	return text
}
```

```go
article

package internal

import "strings"

// FixArticles changes "a" to "an" before vowels or 'h'
func FixArticles(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {

		if words[i] == "a" || words[i] == "A" {

			nextWord := strings.ToLower(words[i+1])

			if startsWithVowelOrH(nextWord) {

				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}

			}
		}
	}

	return strings.Join(words, " ")
}
```

```go
convert.go


package utils

import (
	"strconv"
)

// HexToDecimal converts a hexadecimal string to decimal
func HexToDecimal(hex string) int {

	num, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0
	}

	return int(num)
}

// BinToDecimal converts a binary string to decimal
func BinToDecimal(bin string) int {

	num, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return 0
	}

	return int(num)
}
```

```go
case

package utils

import "strings"

// ToUpper converts a word to uppercase
func ToUpper(word string) string {
	return strings.ToUpper(word)
}

// ToLower converts a word to lowercase
func ToLower(word string) string {
	return strings.ToLower(word)
}

// Capitalize makes the first letter uppercase and the rest lowercase
func Capitalize(word string) string {

	if len(word) == 0 {
		return word
	}

	first := strings.ToUpper(string(word[0]))
	rest := strings.ToLower(word[1:])

	return first + rest
}
```

```go
modifier text

package internal

import "testing"

func TestHexModifier(t *testing.T) {

	input := "1E (hex) files were added"
	expected := "30 files were added"

	result := ApplyModifiers(input)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestBinModifier(t *testing.T) {

	input := "It has been 10 (bin) years"
	expected := "It has been 2 years"

	result := ApplyModifiers(input)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestUpModifier(t *testing.T) {

	input := "Ready set go (up)"
	expected := "Ready set GO"

	result := ApplyModifiers(input)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestLowModifier(t *testing.T) {

	input := "I should stop SHOUTING (low)"
	expected := "I should stop shouting"

	result := ApplyModifiers(input)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestCapModifier(t *testing.T) {

	input := "welcome to the brooklyn bridge (cap)"
	expected := "welcome to the brooklyn Bridge"

	result := ApplyModifiers(input)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
``` 

```go
punctuation test

package internal

import "testing"

func TestFixPunctuation(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
		},
		{
			input:    "Hello , world !",
			expected: "Hello, world!",
		},
		{
			input:    "Wait ... what ?",
			expected: "Wait... what?",
		},
		{
			input:    "I am excited : are you ?",
			expected: "I am excited: are you?",
		},
		{
			input:    "Yes ; absolutely !",
			expected: "Yes; absolutely!",
		},
	}

	for _, test := range tests {
		result := FixPunctuation(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}
```

```go
convert test


package utils

import "testing"

func TestHexToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1E", 30},
		{"FF", 255},
		{"0", 0},
		{"A", 10},
	}

	for _, test := range tests {
		result := HexToDecimal(test.input)
		if result != test.expected {
			t.Errorf("HexToDecimal(%s) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestBinToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"10", 2},
		{"1010", 10},
		{"0", 0},
		{"1", 1},
	}

	for _, test := range tests {
		result := BinToDecimal(test.input)
		if result != test.expected {
			t.Errorf("BinToDecimal(%s) = %d; want %d", test.input, result, test.expected)
		}
	}
}
```
