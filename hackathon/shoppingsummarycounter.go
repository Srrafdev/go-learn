package piscine

func chang(s string) []string {
	var txt []string
	txt = append(txt, "")
	a := 0
	for i := 0; i < len(s); i++ {
		if s[i] == (' ') {
			txt = append(txt, "")
			a++
			continue
		}

		txt[a] += string(s[i])
	}
	return txt
}

func ShoppingSummaryCounter(str string) map[string]int {
	resu := map[string]int{}
	splstr := chang(str)
	for i := 0; i < len(splstr); i++ {
		a := 0
		for j := 0; j < len(splstr); j++ {
			if splstr[i] == splstr[j] {
				a++
			}
		}
		resu[splstr[i]] = a
	}
	return resu
}
