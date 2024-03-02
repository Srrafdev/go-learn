package main //piscine

func Compact(ptr *[]string) int {
	var rb []string
	d := 0
	for _, char := range *ptr {
		if char != "" {
			rb = append(rb, char)
			d++
		}
	}
	*ptr = rb
	return d
}
