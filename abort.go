package main

import "fmt"

func Abort(a, b, c, d, e int) int {
	num := []int{a, b, c, d, e}
	for i := 0; i < len(num)-1; i++ {
		start := num[i]
		if num[i] >= num[i+1] {
			num[i] = num[i+1]
			num[i+1] = start

		}
	}
	return num[len(num)/2]

}

func main() {

	fmt.Println(Abort(2, 3, 8, 5, 7))
}
