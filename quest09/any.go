package main //piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for i := 0; i < len(a); i++ {
		result = f(a[i])

		if result == true {
			return true
		}
	}
	return result
}
