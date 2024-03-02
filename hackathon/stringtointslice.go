package main //piscine

func StringToIntSlice(str string) []int {
	var n []int
	if str == "" {
		return n
	}

	for _, char := range str {
		n = append(n, int(char))
	}
	return n
}
