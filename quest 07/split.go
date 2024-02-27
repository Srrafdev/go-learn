package piscine

func Split(s, sep string) []string {
	stri := []string{}
	stri = append(stri, "")
	a := 0
	a2 := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			stri = append(stri, "")
			a++
			i += len(sep)
		}
		stri[a] += string(s[i])
		a2 = i
	}

	for i := a2; i < len(s); i++ {
		stri[a] += string(s[i])
	}
	return stri
}
