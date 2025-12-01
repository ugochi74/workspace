package main

import "fmt"

func LastWord(s string) string {
	last := ""
	current := ""
	for _, r := range s {
		if r == ' ' {
			if current != "" {
				last = current
				current = ""
			}
		} else {
			current += string(r)
		}
	}
	if current != "" {
		last = current
	}
	if last != "" {
		return last + "\n"
	}
	return ""
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
