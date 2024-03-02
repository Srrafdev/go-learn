package main // piscine

func JumpOver(str string) string {
	vari := ""
	if len(str) < 3 {
		return "\n"
	}
	for i := 2; i < len(str); i += 3 {
		vari += string(str[i])
	}
	vari += "\n"
	return vari
}
