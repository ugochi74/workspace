package main

import (
	"fmt"
)

func CountCharacter(s string, c rune) int {
	if s == "" {
		return 0
	}
	count := 0
	for _, i := range s {
		if i == c {

		}
	}
	return count
}
func main() {
	fmt.Println(CountCharacter("", 'l'))
}
