package main 
import "github.com/01-edu/z01"

func CheckNumber(s string) bool {
	if _, r := range s {
		if (r >= '0' && r <= '9') {
			return true
		}
	  
	}
return false
}