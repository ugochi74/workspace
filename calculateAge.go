package main

import "fmt"

func calculateAge(birthyear, birthmonth, currentyear, currentmonth int) (int, int) {
	year := currentyear - birthyear
	month := currentmonth - birthmonth
	if month < 0 {
		year--
		month += 12
	}
	return year, month
}

func main() {
	var birthyear, birthmonth int
	var currentyear, currentmonth int
	fmt.Print("Enter your birth year: ")
	fmt.Scan(&birthyear)
	fmt.Print("Enter your birth month: ")
	fmt.Scan(&birthmonth)
	fmt.Print("Enter the current year: ")
	fmt.Scan(&currentyear)
	fmt.Print("Enter the current month: ")
	fmt.Scan(&currentmonth)

	year, month := calculateAge(birthyear, birthmonth, currentyear, currentmonth)
	fmt.Printf("Your age is: %d year and %d month(s)\n", year, month)
}
