package main //piscine

func Abort(a, b, c, d, e int) int {
	totl := []int{a, b, c, d, e}
	for i := 0; i < len(totl)-1; i++ {
		for j := i + 1; j < len(totl); j++ {
			if totl[i] > totl[j] {
				totl[i], totl[j] = totl[j], totl[i]
			}
		}
	}
	return totl[2]
}
