package main

import "fmt"

func reverseString(s string) string {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}

func main() {
	fmt.Println(reverseString("A man a plan a canal panama"))
	fmt.Println(reverseString("hello"))
	fmt.Println(reverseString("level"))
	fmt.Println(reverseString("madam"))

}
