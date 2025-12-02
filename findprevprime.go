package main

import "fmt"

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return FindPrevPrime(nb - 1)
		}
	}
	return nb
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
