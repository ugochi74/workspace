package main
import "fmt"

func PrintMemory(arr[10]byte) {
    fullASCII := ""
    for i := 0; i < len(arr); i += 4 {
		
        line := " "
        for j := i; j < i+4 && j < len(arr); j++ {
            line += fmt.Sprintf("%02x", arr[j])
            if j < i+3 && j < len(arr) -1 {
                line += " "
            }
            if arr[j] >= 32 && arr[j] <= 126 {
                fullASCII += string(arr[j])
            } else {
                fullASCII += "."
            }
        }
        fmt.Println(line)
    }
    fmt.Println(fullASCII)
}



func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
