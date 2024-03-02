package piscine

func ToLower(s string) string {
	re := []rune(s)
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			re[i] += 32
		} else {
			continue
		}
	}
	return string(re)
}
