package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1:]

	for _, char := range name {
		if char == "01" || char == "galaxy" || char == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
