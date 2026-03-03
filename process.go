package main

import (
	"io/ioutil"
	"strings"
)

func run(inputFile, outputFile string) error {
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	text := string(content)
	result := processText(text)

	return ioutil.WriteFile(outputFile, []byte(result+"\n"), 0644)
}

func processText(text string) string {
	words := strings.Fields(text)
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		// -------------------------
		// Handle case commands
		// -------------------------
		if strings.HasPrefix(word, "(cap") ||
			strings.HasPrefix(word, "(low") ||
			strings.HasPrefix(word, "(up") {

			command := word

			// If split like "(cap," "3)"
			if !strings.HasSuffix(word, ")") && i+1 < len(words) {
				command = word + words[i+1]
				i++ // skip next part
			}

			switch {
			case strings.HasPrefix(command, "(up"):
				result = applyCase(result, "up", command)
			case strings.HasPrefix(command, "(low"):
				result = applyCase(result, "low", command)
			case strings.HasPrefix(command, "(cap"):
				result = applyCase(result, "cap", command)
			}

			continue
		}

		// -------------------------
		// Handle hex/bin commands
		// -------------------------
		switch word {
		case "(hex)":
			if len(result) > 0 {
				result[len(result)-1] = hexToDec(result[len(result)-1])
			}
		case "(bin)":
			if len(result) > 0 {
				result[len(result)-1] = binToDec(result[len(result)-1])
			}
		default:
			result = append(result, word)
		}
	}

	// -------------------------
	// Post processing
	// -------------------------
	finalText := strings.Join(result, " ")
	finalText = fixPunctuation(finalText)
	finalText = fixQuotes(finalText)
	finalText = fixArticles(finalText)

	return finalText
}
