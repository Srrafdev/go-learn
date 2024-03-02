package main // piscine

func CollatzCountdown(start int) int {
	r := 0
	if start <= 0 {
		return -1
	}
	for start != 1 {
		if start%2 == 0 {
			start = start / 2
			r++
		} else if start%2 != 0 {
			start = (start * 3) + 1
			r++
		}
	}
	return r
}
