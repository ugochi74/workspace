package main

import "fmt"

func recursiveSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + recursiveSum(numbers[1:])

}
func main() {
	var res = []int{-2, -2, -3, 4, 5, 6}
	fmt.Println(recursiveSum(res))

}
