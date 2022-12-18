// -- Summary:
// Small program that have passenger tickets and a bus,
// - the bus will indicate who is on the front seat.
// - the passenger ticket will wheter or not a passenger boarded the bus.

package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false} // Create & assign shorthand
	fmt.Println(casey)

  // Block assignment
	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)

  fmt.Println(bill, ella)

  // Un-initialize variable, the variable takes
  // on every field the default values of its type
  var heidi Passenger

  // Assigning values
  heidi.Name = "Heidi"
  heidi.TicketNumber = 4
  fmt.Println(heidi)

  // Boarding passengers
  casey.Boarded = true
  bill.Boarded = true

  if casey.Boarded {
    fmt.Println("Casey has boarded the bus")
  }

  if bill.Boarded {
    fmt.Println(bill.Name,"has boarded the bus")
  }

  heidi.Boarded = true

  bus := Bus{heidi}

  fmt.Println(bus)
  fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
