
func HashCode(dec string) string {
    result := ""
    for _, ch := range dec {
        hash := (int(ch) + len(dec))  %127
        if hash < 33 {
            hash += 33
        }
        result += string(rune(hash))
    }
         return result
}

