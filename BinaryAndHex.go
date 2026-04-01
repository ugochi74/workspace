package main
import (
	"fmt"
	"strconv"
	"strings"
)
func main() {
	var choice string
	for {
		fmt.Println("Choose your choice(convertHex <1>, convertBin <2>, convertDec <3>): ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var input string
			fmt.Println("Enter input: ")
			fmt.Scanln(&input)
			result, err := ParseDecimal(input, 16)
			if err != nil {
				fmt.Println("Invalid input: ")
				continue
			}
			decimal := strconv.FormatInt(result, 10)
			fmt.Printf("decimal: %s\n", strings.ToUpper(decimal))

		case "2":
			var input string
			fmt.Println("Enter input: ")
			fmt.Scanln(&input)
			result, err := ParseDecimal(input, 2)
			if err != nil {
				fmt.Println("Invalid input: ")
				continue
			}
			decimal := strconv.FormatInt(result, 10)
			fmt.Printf("decimal: %s\n", decimal)

		case "3":
			var input string
			fmt.Println("Enter input: ")
			fmt.Scanln(&input)
			result, err := ParseDecimal(input, 10)
			if err!=nil {
				fmt.Println("Invalid input: ")
				continue
			}
			//fmt.Println(result)
			bin := strconv.FormatInt(result, 2)
		hex := strconv.FormatInt(result,16)
		fmt.Printf("hex: %s\n", strings.ToUpper(hex))
		fmt.Printf("bin: %s\n", bin)
        }
	}
}

func ParseDecimal(s string, base int) (int64, error) {
	result, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
