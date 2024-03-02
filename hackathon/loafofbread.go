package piscine

func LoafOfBread(str string) string {
	restr, idx := "", 0

	if len(str) < 5 {
		return "Invalid Output\n"
	}

	for _, char := range str {
		if char == ' ' {
			continue
		}
		if idx%6 == 5 {
			restr += " "
		} else {
			restr += string(char)
		}
		idx++
	}

	return restr + "\n"

}
