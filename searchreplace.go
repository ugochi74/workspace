package piscine

func SearchReplace(s, old, new string) string {
	if len(old) != 1 || len(new) != 1 {
		return s
	}
	o := old[0]
	n := new[0]
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == o {
			result += string(n)
		} else {
			result += string(s[i])
		}
	}

	return result
}
