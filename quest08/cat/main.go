package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printz01(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	var fileargs string
	cron := 0
	if len(os.Args) == 1 {
		data, _ := io.ReadAll(os.Stdin)
		printz01(string(data))
	}
	for i := 1; i < len(os.Args); i++ {
		cron++

		fileargs = os.Args[i]

		file, err := os.Open(fileargs)
		if err != nil {
			printz01("ERROR: open " + fileargs + ": no such file or directory")
			z01.PrintRune('\n')
			os.Exit(1)
		}

		data, _ := os.ReadFile(fileargs)

		file.Close()
		printz01(string(data))
	}

	// data, err := io.ReadAll(os.Stdin)

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "ERROR: from stdin: %v\n", err)
	// 	os.Exit(1)
	// }

	// printz01(string(data))
}
