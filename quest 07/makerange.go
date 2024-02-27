package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	if min == 0 || max == 0 {
		return []int{0}
	}

	sli := make([]int, max-min)
	for i := min; i < max; i++ {
		sli[i-min] = i
	}
	return sli
}
