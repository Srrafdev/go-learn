package main

import "fmt"

func Atoi(s string) int {
	idx := 0
	nigativ := 1
	x := 0

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

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
