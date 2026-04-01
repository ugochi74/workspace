package main

import "fmt"

func main() {
	// var input string
	var distance float64
	var choice int
	for {
		fmt.Println("distance converter(kilometer to meter <1>, meters to kilometer <2>): ")
		//fmt.Println("1. kilometer to meter")
		//fmt.Println("2. meters to kilometer")
		fmt.Print("Choose choice: ")
		fmt.Scanln(&choice)
		fmt.Println("Enter distance")
		fmt.Scanln(&distance)

		switch choice {
		case 1:
			result := distance * 1000
			fmt.Printf(" %.2f km = %.2f m\n", distance, result)
			continue

		case 2:
			result := distance / 1000
			fmt.Printf(" %.2f m = %.2f km\n", distance, result)

			fmt.Printf(" %.3f m = %.3f km\n", distance, result)
			continue
		default:
			fmt.Println("invalid choice!")
			continue

		}

	}
}
