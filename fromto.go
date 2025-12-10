package main
import "fmt"

func FromTo(from, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}
	if from == to {
		return fmt.Sprintf("%02d\n", from)
	}
	out := ""
	step := 1
	if from > to {
		step = -1
	}
	for i := from; i != to+step; i += step {
		if i != from {
			out += ","
		}
		out += fmt.Sprintf("%02d", i)
	}
	return out + "\n"
}
