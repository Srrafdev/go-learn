package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	var s string

	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	sli := []rune(s)

	for i := 0; i < len(sli)-1; i++ {
		for j := i + 1; j < len(sli); j++ {
			if sli[i] > sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}

	for _, d := range sli {
		z01.PrintRune(rune(d))
	}
}
