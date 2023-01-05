// --Summary:
// A program that can track the status of parking spaces

package main

import "fmt"

// Represents a single parking space
type Space struct {
	occupied bool
}

// Represents all the sapces that we have
type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNumber int) {
	lot.spaces[spaceNumber-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNumber int) {
	lot.spaces[spaceNumber-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNumber int) {
	lot.spaces[spaceNumber-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 5)}
	fmt.Println("Initial:", lot)

	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupied", lot)

	lot.vacateSpace(2)
	fmt.Println("After vacate", lot)
}
