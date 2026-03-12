package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func hexToDecimal(hexStr string) (int64, error) {
	dec, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return 0, err
	}
	return dec, nil
}

func binToDecimal(binStr string) (int64, error) {
	dec, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return 0, err
	}
	return dec, nil
}

func capitalizeWord(word string) string {
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])

	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func toLower(word string) string {
	return strings.ToLower(word)
}

func toUpper(word string) string {
	return strings.ToUpper(word)
}
func main() {
	fmt.Println(hexToDecimal("1E")) // prints 30
	fmt.Println(hexToDecimal("FF")) // prints 255
	fmt.Println(binToDecimal("10"))
	fmt.Println(binToDecimal("1010"))
	fmt.Println(binToDecimal("11111111"))
	fmt.Println(capitalizeWord("hELLO"))
	fmt.Println(capitalizeWord("WORLD"))
	fmt.Println(toLower("hELLO"))
	fmt.Println(toUpper("world"))
}






package main

import (
    "fmt"
    "strings"
    
    )
    
func Capitalize(word []string, n int) []string {
    //modified := strings.ReplaceAll(s, " ,", ",")
    //text := strings.ToLower(s)
    //return strings.Title(text)
   // word = strings.ToLower(word[:(len(word)-1)])
   for i := 0; i < len(word) && i < n; i++  {
       word[i] = strings.ToUpper(string(word[i][:1])) + strings.ToLower(string(word[i][1:]))
   }
  // text := strings.Join(word, " ")
   return word
}

func main() {
    fmt.Println(AToAn("A apple!"))
}

func toLower(word string) string {
    return strings.ToLower(word)
}

func AToAn(s string) string {
    text := strings.Fields(s)
    for i := 0; i < len(text)-1; i++ {
    first := text[i+1]
    slice := first[0]
    boolean := strings.ContainsAny(string(slice), "aeiouhAEIOUH")
    if text[i] == "a" && boolean == true {
        text[i] = "an"
    } else if text[i] == "A" && boolean == true{
        text[i] = "An"
    }
    }
   str := strings.Join(text, " ")
    return str
}







package main

import (
	"fmt"
	"strconv"
	"strings"
)

func capitalizePreviousN(word []string, n int) []string {
	for i := 0; i < len(word) && i < n; i++ {
		if i == 2 {
			continue
		}

		word[i] = strings.ToUpper(word[i])
	}
	return word

}

func main() {
	var res = []string{"i", "am", "excited", "confidence"}
	fmt.Println(capitalizePreviousN(res, 4))

}

func capitalize(s string) string {
	text := strings.ToUpper(s)
	return strings.Title(text)
    strings.ContainsAny(s, chars)
}

func hexToDecimal(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}






package main

import (
	"fmt"
	"strings"
)

func capitalizePreviousN(word []string, n int) []string {
	for i := 0; i < len(word); i++ {
		if i < n {
			word[i] = strings.ToUpper(string(word[i][:1])) + strings.ToLower(string(word[i][1:]))
		} else {
			word[i] = strings.ToLower(string(word[i]))
		}
	}
	return word
}

func main() {
	var res = []string{"hEllo", "how", "ARE", "yOu"}
	fmt.Println(capitalizePreviousN(res, 3))
}


