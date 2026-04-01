package main

import (
	"fmt"
	"strings"
)

func main() {
	var input, operation string
	for {
		fmt.Println("Enter input: ")
		fmt.Scanln(&input)

		fmt.Println("Choose operation(lastChar, capitalize, firstlast, lastindex, deleteindex): ")
		fmt.Scanln(&operation)

		switch operation {
		case "lastChar":
			fmt.Println(string(input[len(input)-1]))
		case "capitalize":
			fmt.Println(strings.ToUpper(input))
		case "firstlast":
			if len(input) == 0 {
				fmt.Println(input)
			} else {
				first := string(input[0])
				last := string(input[len(input)-1])
				fmt.Println(first, last)
			}
		case "lastindex":
			fmt.Println(input[:len(input)-1] + strings.ToUpper(input[len(input)-1:]))
		case "deleteindex":
			var index int
			fmt.Println("enter index to delete: ")
			fmt.Scanln(&index)

			if index < 0 || index > len(input) {
				fmt.Println("Error invalid index: ")
				continue
			}
			fmt.Println(input[:index] + input[index+1:])
		default:
			fmt.Println("Invalid operation: ")
			continue
		}
		break
	}

}
