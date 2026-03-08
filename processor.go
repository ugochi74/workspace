package main

func Process(text string) string {

	text = ConvertNumbers(text)

	text = ApplyCase(text)

	text = FixPunctuation(text)

	text = FixArticles(text)

	return text
}