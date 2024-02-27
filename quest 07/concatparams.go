package piscine

func ConcatParams(args []string) string {
	var stri string
	for i := 0; i < len(args); i++ {
		stri += args[i]
		if i != len(args)-1 {
			stri += "\n"
		}
	}
	return stri
}
