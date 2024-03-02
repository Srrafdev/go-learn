package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}

	file, _ := os.Open(os.Args[1])
	arr := make([]byte, 14)
	file.Read(arr)

	if arr != nil {
		fmt.Println(string(arr))
		return
	}

	file.Close()
}
