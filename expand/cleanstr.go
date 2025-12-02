package main

import "os"

func main() {
	if len(os.Args) != 2 {
		os.Stdout.Write([]byte("\n"))
		return
	}
	s, out, space := os.Args[1], "", false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' && c != '\t' {
			out += string(c)
			space = true
		} else if space {
			out += "  "
			space = false
		}
	}
	if len(out) > 0 && out[len(out)-1] == ' ' {
		out = out[:len(out)-1]
	}
	if out == "" {
		os.Stdout.Write([]byte("\n"))
	} else {
		os.Stdout.Write([]byte(out + "\n"))
	}
}
