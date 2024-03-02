package main // piscine

func Join(strs []string, sep string) string {
	x := ""
	for idx, char := range strs {
		x += char
		if idx == len(strs)-1 {
			break
		}
		x += sep
	}
	return x
}
