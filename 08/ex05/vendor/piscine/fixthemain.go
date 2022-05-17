package piscine

import "ft"

type Door struct {
	state int
}

const OPEN = 0
const CLOSE = 1

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	return Door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

