package piscine

func ActiveBits(n int) int {
	in := 0
	b := 0
	for n != 0 {
		in = n % 2
		n /= 2
		if in == 1 {
			b++
		}
	}

	return b
}
