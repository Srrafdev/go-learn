package main // piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	is := true
	im := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			is = false
		}
	}
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			im = false
		}
	}
	if im {
		return true
	}
	return is
}

func is(a, b int) int {
	if a < b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
