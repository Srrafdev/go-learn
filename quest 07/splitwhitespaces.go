package piscine

func SplitWhiteSpaces(s string) []string {
	var stri []string
	stri = append(stri, "")
	x := 0
	for i := 0; i < len(s); i++ {
		if s[i] == (' ') || s[i] == ('\t') || s[i] == ('\n') {
			if s[i+1] == ' ' || s[i+1] == ('\t') || s[i+1] == ('\n') {
				continue
			}
			stri = append(stri, "")
			x++
			continue
		}

		stri[x] += string(s[i])
	}
	return stri
}
