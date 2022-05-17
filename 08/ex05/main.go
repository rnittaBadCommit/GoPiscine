package main

import (
	. "piscine"
)

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.State == OPEN {
		CloseDoor(door)
	}
}
