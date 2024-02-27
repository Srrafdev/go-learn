package piscine

func AppendRange(min, max int) []int {
	if min == 0 && max == 0 || min >= max {
		return nil
	}

	if min == 0 || max == 0 {
		return []int{0}
	}

	var sli []int
	for i := min; i < max; i++ {
		sli = append(sli, i)
	}
	return sli
}
