package main

import "fmt"

func PrintIf(s string) string {
	if s == "" {
		return "G\n"
	}
	if len(s) >= 3 {
		return "G\n"
	}
	return "invalid input\n"
}
func main() {
	fmt.Print(PrintIf("qwer2d"))
	fmt.Print(PrintIf("afd"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("13"))
}
