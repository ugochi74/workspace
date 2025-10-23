package main

import "github.com/01-edu/z01"

func CountAlpha(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count++
		}
	}
	return count
}

func main() {
	result := CountAlpha("Hello World 123!")
	z01.PrintRune('0' + rune(result/10))
	z01.PrintRune('0' + rune(result%10))
	z01.PrintRune('\n')
}
