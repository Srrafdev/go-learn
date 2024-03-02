package piscine

func IsLower(s string) bool {
	isuper := false
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			isuper = true
		} else {
			return false
		}
	}
	return isuper
}
