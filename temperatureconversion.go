package main

import "fmt"

func main() {
	var choice int
	var temperature float64
	for {
		fmt.Println("Temperature to convert(celeius to fahrenheit<1>, fahrenheit to celeius <2>):")
		fmt.Println("choose choice: ")
		fmt.Scanln(&choice)
		fmt.Println("Enter temperature: ")
		fmt.Scanln(&temperature)
		switch choice {
		case 1:
			result := (temperature * 9 / 5) + 32
			fmt.Printf("%.2f %cC = %.2f %cF\n", temperature, 176, result, 176)
			continue
		case 2:
			result := (temperature - 32) * 5 / 9
			fmt.Printf("%.2f F = %.2f C\n", temperature, result)
			continue
		default:
			fmt.Println("Invalid temperature: ")
			continue
		}
	}
}
