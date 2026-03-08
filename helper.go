package main

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) (string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func WriteFile(filename string, content string) error {

	return os.WriteFile(filename, []byte(content), 0644)
}

func ConvertNumbers(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		if words[i] == "(hex)" && i > 0 {

			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}

		if words[i] == "(bin)" && i > 0 {

			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

func ApplyCase(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		// Handle (up) (low) (cap)
		if words[i] == "(up)" || words[i] == "(low)" || words[i] == "(cap)" {

			if i > 0 {
				applyCase(&words[i-1], words[i])
			}

			// remove modifier
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		// Handle numbered modifiers (up, 2) (low, 3) (cap, 4)
		if strings.HasPrefix(words[i], "(up,") ||
			strings.HasPrefix(words[i], "(low,") ||
			strings.HasPrefix(words[i], "(cap,") {

			if i+1 < len(words) {

				// extract command
				command := strings.TrimPrefix(words[i], "(")
				command = strings.TrimSuffix(command, ",")

				// extract number
				number := strings.TrimSuffix(words[i+1], ")")

				n, err := strconv.Atoi(number)
				if err == nil {

					for j := 1; j <= n && i-j >= 0; j++ {
						applyCase(&words[i-j], "("+command+")")
					}
				}

				// remove "(up," and "2)"
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}
	}

	return strings.Join(words, " ")
}

func applyCase(word *string, command string) {

	switch command {

	case "(up)":
		*word = strings.ToUpper(*word)

	case "(low)":
		*word = strings.ToLower(*word)

	case "(cap)":
		if len(*word) > 0 {
			w := *word
			*word = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
		}
	}
}

func FixPunctuation(text string) string {

	text = strings.ReplaceAll(text, " ,", ",")
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ;", ";")

	return text
}

func FixArticles(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {

		if strings.ToLower(words[i]) == "a" {

			next := strings.ToLower(words[i+1])

			if strings.HasPrefix(next, "a") ||
				strings.HasPrefix(next, "e") ||
				strings.HasPrefix(next, "i") ||
				strings.HasPrefix(next, "o") ||
				strings.HasPrefix(next, "u") ||
				strings.HasPrefix(next, "h") {

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
