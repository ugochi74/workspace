package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	prevUpper := false
	for i, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') {
			return s
		}
		if r >= 'A' && r <= 'Z' {
			if prevUpper || i == len(s)-1 {
				return s
			}
			prevUpper = true
		} else {
			prevUpper = false
		}
	}
	result := ""
	for i, r := range s {
		if r >= 'A' && r <= 'Z' && i != 0 {
			result += "_"
		}
		result += string(r)
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
