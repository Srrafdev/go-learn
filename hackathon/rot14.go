package piscine

func Rot14(s string) string {
	str := []byte(s)
	for i := 0; i <= len(s)-1; i++ {
		if str[i] >= 'a' && str[i] < 'm' || str[i] >= 'A' && str[i] < 'M' {
			str[i] += 14
		} else if str[i] >= 'm' && str[i] <= 'z' || str[i] >= 'M' && str[i] <= 'Z' {
			str[i] -= 12
		} else {
			continue
		}
	}
	return string(str)
}
