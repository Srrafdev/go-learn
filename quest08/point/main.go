package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

const s string = "x = 42, y = 21\n"

func main() {
	points := &point{}

	setPoint(points)

	for _, char := range s {
		z01.PrintRune(char)
	}
	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
