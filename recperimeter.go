package main
import "fmt"

func RecPerimeter(w,h int) int {
	if w < 0 || h < 0 {
		return -1
	}
       return 2* (w*h)
}

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}
