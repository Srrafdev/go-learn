package piscine

func IsUpper(s string) bool {
	isuper := false
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			isuper = true
		} else {
			return false
		}
	}
	return isuper
}
