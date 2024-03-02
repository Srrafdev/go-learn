package piscine

func Max(a []int) int {
	max := a[0]
	for _, char := range a {
		if max < char {
			max = char
		}
	}
	return max
}
