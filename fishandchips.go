
func FishAndChips(n int) string {
    if n < 0 {
        return "error"
    }
    if n%2 == 0 || n%3 == 0 {
        return "fish, chips"
    }
    if n%2 == 0 {
        return "fish"
    }
    if n%3 == 0 {
        return "chips"
    } else {
        return "error"
    }
}
