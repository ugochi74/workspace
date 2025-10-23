package main

func RetainFirstHalf(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return s
	}
	midpoint := length / 2
	return s[:midpoint]
}
