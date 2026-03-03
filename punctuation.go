package main

import (
	"strings"
)

// Fix spacing around punctuation marks
func fixPunctuation(text string) string {
	// Handle special groups first (like ... or !?)
	text = strings.ReplaceAll(text, " . . .", "...")
	text = strings.ReplaceAll(text, " ...", "...")
	text = strings.ReplaceAll(text, "!? ", "!? ")

	// Standard punctuation marks
	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuations {
		// Remove space before punctuation
		text = strings.ReplaceAll(text, " "+p, p)
		// Ensure space after punctuation (except at end of text)
		text = strings.ReplaceAll(text, p+" ", p+" ")
	}

	return text
}

// Fix quotes: ensure 'word' formatting
func fixQuotes(text string) string {
	words := strings.Fields(text)
	var result []string
	insideQuotes := false
	var quotedPhrase string

	for _, w := range words {
		if w == "'" {
			if insideQuotes {
				// Closing quote
				quotedPhrase += "'"
				result = append(result, quotedPhrase)
				insideQuotes = false
				quotedPhrase = ""
			} else {
				// Opening quote
				insideQuotes = true
				quotedPhrase = "'"
			}
		} else {
			if insideQuotes {
				// Add word to the quoted phrase
				if quotedPhrase == "'" {
					quotedPhrase += w
				} else {
					quotedPhrase += " " + w
				}
			} else {
				// Normal word
				result = append(result, w)
			}
		}
	}

	return strings.Join(result, " ")
}
