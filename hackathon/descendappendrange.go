package main // piscine

func DescendAppendRange(max, min int) []int {
	if min >= max {
		return []int{}
	}
	var sli []int
	for i := max; i > min; i-- {
		sli = append(sli, i)
	}
	return sli
}
