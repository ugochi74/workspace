package main

import "fmt"

func FirstWord(s string) string {
	started := false
	result := ""

	if len(s) == 0 {
		return "\n"
	}

	for _, r := range s {
		if r == ' ' {
			if started {
				break
			}
			continue
		}
		started = true
		result += string(r)
	}

	if result != "" {
		result += "\n"
	}

	return result
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
