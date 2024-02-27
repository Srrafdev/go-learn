package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	stri := ""
	for i := 0; i < len(a); i++ {
		stri = a[i] + "\n"
		for _, char := range stri {
			z01.PrintRune(char)
		}
	}
}
