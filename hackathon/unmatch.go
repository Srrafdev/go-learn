package main // piscine

func Unmatch(a []int) int {
	result := 0
	for _, char := range a {
		for _, char2 := range a {
			if char == char2 {
				result++
			}
		}
		if result%2 != 0 {
			return char
		}
	}
	return -1
}
