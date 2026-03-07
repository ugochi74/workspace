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

		if words[i] == "(up)" && i > 0 {

			words[i-1] = strings.ToUpper(words[i-1])

			words = append(words[:i], words[i+1:]...)
			i--
		}

		if words[i] == "(low)" && i > 0 {

			words[i-1] = strings.ToLower(words[i-1])

			words = append(words[:i], words[i+1:]...)
			i--
		}

		if words[i] == "(cap)" && i > 0 {

			word := words[i-1]

			if len(word) > 0 {
				words[i-1] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
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
