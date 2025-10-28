package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Check if valid camelCase or PascalCase
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
			return s // contains invalid char
		}
		if i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z' && ch >= 'A' && ch <= 'Z' {
			return s // two consecutive uppercase
		}
	}
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s // ends with uppercase
	}

	// Convert to snake_case
	res := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' && i != 0 {
			res += "_"
		}
		res += string(ch)
	}
	return res
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
