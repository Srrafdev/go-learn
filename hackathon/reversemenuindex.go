package piscine

func ReverseMenuIndex(menu []string) []string {
	s := make([]string, len(menu))
	rach := 0
	for i := len(menu) - 1; i >= 0; i-- {
		s[rach] = menu[i]
		rach++
	}
	return s
}
