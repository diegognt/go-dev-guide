// - Summary:
// This demo program use a Room Go struct to represent a hotel room
// and its cleanliness. There is a function that iterate over an array of rooms
// and print whether the a room is clean or not.

package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
  rooms := [...]Room {
    {name: "Office"},
    {name: "Warehouse"},
    {name: "Reception"},
    {name: "Ops"},
  }

  checkCleanliness(rooms)

  fmt.Println("Performing cleaning...")

  rooms[1].cleaned = true
  rooms[3].cleaned = true

  checkCleanliness(rooms)

}
