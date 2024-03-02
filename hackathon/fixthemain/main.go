package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

type door struct {
	state string
}

func OpenDoor(Door *door) {
	PrintStr("Door Opening...")
	Door.state = "OPEN"
}

func CloseDoor(Door *door) {
	PrintStr("Door Closing...")
	Door.state = "CLOSE"
}

func IsDoorOpen(Door *door) bool {
	PrintStr("is the Door opened ?")
	if Door.state == "OPEN" {
		return true
	} else {
		return false
	}
}

func IsDoorClose(Door *door) bool {
	PrintStr("is the Door closed ?")
	if Door.state == "CLOSE" {
		return true
	} else {
		return false
	}
}

func main() {
	door := &door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
}
