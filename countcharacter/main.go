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
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountCharacter("Hello World", 'l'))
	fmt.Println(CountCharacter("5  balloons", 5))
	fmt.Println(CountCharacter("   ", ' '))
	fmt.Println(CountCharacter("The 7 deadly sins", '7'))
}
