package piscine

func AlphaCount(s string) int {
	var result int
	for _, char := range s {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			result++
		}
	}
	return result
}
