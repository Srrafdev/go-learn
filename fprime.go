package main

import (
	"fmt"
)

// import "fmt"

func main() {
	// fp := 0
	str := 42
	// ray := []int{}

	for i := 2; i <= str; i++ {
		if IsPrime(i) {
			if (str % i) == 0 {
				fmt.Print(i)
				fmt.Print("*")
				str /= i
				i--
			}
		}
	}
	 fmt.Print("\n")
	 
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
	}
	return true
}

func Atoi(s string) int {
	idx := 0
	nigativ := 1
	x := 0

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			nigativ *= -1
		}
		x++
	}

	for i := x; i < len(s); i++ {
		if s[x] >= '0' && s[i] <= '9' {
			idx *= 10
			idx += int(s[x]) - 48
			x++
		} else {
			return 0
		}
	}
	return idx * nigativ
}
