package main

import "fmt"

func TrimAtoi(s string) int {
	str := ""
	sign := 1

	for _, char := range s {
		if char >= '0' && char <= '9' {
			str += string(char)
		} else if char == '-' && str == "" {
			sign = -1
		}
	}

	result := 0
	multi := 1

	for i := len(str) - 1; i >= 0; i-- {
		digit := int(str[i]) - '0'
		result += digit * multi
		multi *= 10
	}

	return result * sign
}

func main() {
	fmt.Println(TrimAtoi("185345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
