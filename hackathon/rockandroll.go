package piscine

func RockAndRoll(n int) string {
	resul := ""
	if n < 0 {
		return "error: number is negative" + "\n"
	}
	if n%2 == 0 && n%3 == 0 {
		resul = "rock and roll" + "\n"
	} else if n%2 == 0 {
		resul = "rock" + "\n"
	} else if n%3 == 0 {
		resul = "roll" + "\n"
	} else {
		resul = "error: non divisible" + "\n"
	}
	return resul
}
