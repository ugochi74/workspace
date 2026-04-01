package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string
	var result float64
	for {
		fmt.Println("Enter num1: ")
		fmt.Scanln(&num1)

		fmt.Println("Enter operator(+, *, %, /, -): ")
		fmt.Scanln(&operator)

		fmt.Println("Enter num2: ")
		fmt.Scanln(&num2)

		switch operator {
		case "*":
			result = num1 * num2
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "/":
			if num2 == 0 {
				fmt.Println("No division by 0: ")
				return
			}
			result = num1 / num2
		case "%":
			if num2 == 0 {
				fmt.Println("No mod by 0")
			}
			result = float64(int(num1) % int(num2))
		default:
			fmt.Println("invalid operation")
			return
		}
		fmt.Println(result)
		var options string
		fmt.Println("choose options(exit, continue): ")
		fmt.Scanln(&options)
		if options == "exit" {
			break
		}

	}
}
