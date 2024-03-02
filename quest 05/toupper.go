package piscine

func ToUpper(s string) string {
	re := []rune(s)
	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			re[i] -= 32
		} else {
			continue
		}
	}
	return string(re)
}
