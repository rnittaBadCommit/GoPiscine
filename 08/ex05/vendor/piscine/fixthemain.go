package piscine

import "ft"

type Door struct {
	State int
}

const OPEN = 0
const CLOSE = 1

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.State = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.State = CLOSE
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	return Door.State == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.State == CLOSE
}

