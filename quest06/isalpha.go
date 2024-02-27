package piscine

func IsAlpha(s string) bool {
	isuper := false
	for _, char := range s {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
			isuper = true
		} else {
			return false
		}
	}
	return isuper
}
